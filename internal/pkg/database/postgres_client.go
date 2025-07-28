package database

import (
	"beta-book-api/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type PostgresClient struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
	dbSSLMode  string
	logger     zerolog.Logger
}

func NewPostgresClient(cfg *config.AppConfig, logger zerolog.Logger) PostgresClient {
	return PostgresClient{
		dbHost:     cfg.DBHost,
		dbPort:     cfg.DBPort,
		dbUser:     cfg.DBUser,
		dbPassword: cfg.DBPassword,
		dbName:     cfg.DBName,
		dbSSLMode:  cfg.DBSSLMode,
		logger:     logger,
	}
}

func (p *PostgresClient) InitPostgresDB() *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.dbHost, p.dbPort, p.dbUser, p.dbPassword, p.dbName, p.dbSSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		p.logger.Fatal().Err(err).Msgf("❌ Failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		p.logger.Fatal().Err(err).Msgf("❌ Failed to ping DB: %v", err)
	}
	p.logger.Info().Msgf("✅ Connected to PostgreSQL successfully")
	return db
}
