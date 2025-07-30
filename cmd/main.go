// @title           Beta Book API
// @version         1.0
// @description     API service to manage books using Clean Architecture

// @contact.name   ADF Code
// @contact.url    https://github.com/adf-code

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Using token header using the Bearer scheme. Example: "Bearer {token}"

package main

import (
	"beta-book-api/config"
	_ "beta-book-api/docs"
	deliveryHttp "beta-book-api/internal/delivery/http"
	pkgDatabase "beta-book-api/internal/pkg/database"
	pkgLogger "beta-book-api/internal/pkg/logger"
	pkgEmail "beta-book-api/internal/pkg/mail"
	pkgOS "beta-book-api/internal/pkg/object_storage"
	"beta-book-api/internal/repository"
	"beta-book-api/internal/usecase"
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	_ = godotenv.Load() // Load .env

	// Load env config
	cfg := config.LoadConfig()

	logger := pkgLogger.InitLoggerWithTelemetry(cfg)
	sendGridClient := pkgEmail.NewSendGridClient(cfg, logger)
	mail := sendGridClient.InitSendGrid()
	postgresClient := pkgDatabase.NewPostgresClient(cfg, logger)
	db := postgresClient.InitPostgresDB()
	minioClient := pkgOS.NewMinioClient(cfg, logger)
	objectStorage := minioClient.InitMinio()

	// Repository and HTTP handler
	bookRepo := repository.NewBookRepo(db)
	bookCoverRepo := repository.NewBookCoverRepo(db)
	bookUC := usecase.NewBookUseCase(bookRepo, db, logger, mail)
	bookCoverUC := usecase.NewBookCoverUseCase(bookCoverRepo, db, logger, objectStorage)
	handler := deliveryHttp.SetupHandler(bookUC, bookCoverUC, logger)

	// HTTP server config
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: handler,
	}

	// Run server in goroutine
	go func() {
		logger.Info().Msgf("🟢 Server running on http://localhost:%s", cfg.Port)
		logger.Info().Msgf("📚 Swagger running on http://localhost:%s/swagger/index.html", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msgf("❌ Server failed: %v", err)
		}
	}()

	// Setup signal listener
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info().Msgf("🛑 Gracefully shutting down server...")

	// Graceful shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown HTTP server
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msgf("❌ Server shutdown failed: %v", err)
	}

	// ✅ Close PostgreSQL DB
	closePostgres(db, logger)

	logger.Info().Msgf("✅ Server shutdown completed.")
}

func closePostgres(db *sql.DB, logger zerolog.Logger) {
	if err := db.Close(); err != nil {
		logger.Info().Msgf("⚠️ Failed to close PostgreSQL connection: %v", err)
	} else {
		logger.Info().Msgf("🔒 PostgreSQL connection closed.")
	}
}
