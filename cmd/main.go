package main

import (
	"beta-book-api/config"
	deliveryHttp "beta-book-api/internal/delivery/http"
	pkgDatabase "beta-book-api/internal/pkg/database"
	pkgLogger "beta-book-api/internal/pkg/logger"
	pkgEmail "beta-book-api/internal/pkg/mail"
	"beta-book-api/internal/repository"
	"beta-book-api/internal/usecase"
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	emailClient := pkgEmail.NewSendGridClient(cfg, logger)
	postgresClient := pkgDatabase.NewPostgresClient(cfg, logger)
	db := postgresClient.InitPostgresDB()

	// Repository and HTTP handler
	repo := repository.NewBookRepo(db)
	bookUC := usecase.NewBookUseCase(repo, logger)
	handler := deliveryHttp.SetupHandler(bookUC, logger, emailClient)

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
			log.Fatal().Err(err).Msgf("❌ Server failed: %v", err)
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

func closePostgres(db *sql.DB, telemetryLog zerolog.Logger) {
	if err := db.Close(); err != nil {
		telemetryLog.Info().Msgf("⚠️ Failed to close PostgreSQL connection: %v", err)
	} else {
		telemetryLog.Info().Msgf("🔒 PostgreSQL connection closed.")
	}
}
