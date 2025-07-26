package main

import (
	"beta-book-api/config"
	deliveryHttp "beta-book-api/internal/delivery/http"
	"beta-book-api/internal/repository"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load env config
	cfg := config.LoadConfig()

	// Init DB
	db := config.InitPostgresDB(cfg)

	// ✅ Ping to test DB connection
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	} else {
		log.Println("✅ Connected to PostgreSQL successfully")
	}

	// Repository and HTTP handler
	repo := repository.NewBookRepo(db)
	handler := deliveryHttp.SetupHandler(repo)

	// HTTP server config
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: handler,
	}

	// Run server in goroutine
	go func() {
		log.Printf("🟢 Server running on http://localhost:%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed: %v", err)
		}
	}()

	// Setup signal listener
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Gracefully shutting down server...")

	// Graceful shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown HTTP server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server shutdown failed: %v", err)
	}

	// ✅ Close PostgreSQL DB
	closePostgres(db)

	log.Println("✅ Server shutdown completed.")
}

func closePostgres(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("⚠️ Failed to close PostgreSQL connection: %v", err)
	} else {
		log.Println("🔒 PostgreSQL connection closed.")
	}
}
