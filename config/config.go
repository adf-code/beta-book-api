package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	Port                string
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBSSLMode           string
	Env                 string
	TelemetryAPIKey     string
	TelemetryEndpoint   string
	SendGridAPIKey      string
	SendGridSenderEmail string
}

func LoadConfig() *AppConfig {
	// Load from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	return &AppConfig{
		Env:                 getEnv("ENV", "development"),
		Port:                getEnv("APP_PORT", "8080"),
		DBHost:              getEnv("DB_HOST", "localhost"),
		DBPort:              getEnv("DB_PORT", "5432"),
		DBUser:              getEnv("DB_USER", "postgres"),
		DBPassword:          getEnv("DB_PASSWORD", ""),
		DBName:              getEnv("DB_NAME", "bookdb"),
		DBSSLMode:           getEnv("DB_SSLMODE", "disable"),
		TelemetryAPIKey:     getEnv("TELEMETRY_API_KEY", "not_set"),
		TelemetryEndpoint:   getEnv("TELEMETRY_ENDPOINT", "not_set"),
		SendGridAPIKey:      getEnv("SENDGRID_API_KEY", "not_set"),
		SendGridSenderEmail: getEnv("SENDGRID_SENDER_EMAIL", "not_set"),
	}
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
