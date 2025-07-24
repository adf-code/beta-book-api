package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitPostgresDB(cfg *AppConfig) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping DB: %v", err)
	}
	return db
}
