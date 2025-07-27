package main

import (
	"beta-book-api/config"
	"beta-book-api/internal/migration"
	"beta-book-api/internal/pkg/database"
	"log"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate.go [up|down]")
	}
	db := database.InitPostgresDB(cfg)
	dir := "migration" // folder sql

	switch os.Args[1] {
	case "up":
		migration.MigrateUp(db, dir)
	case "down":
		migration.MigrateDown(db, dir)
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}
}
