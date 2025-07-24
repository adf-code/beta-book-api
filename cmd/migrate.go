package main

import (
	"beta-book-api/config"
	"beta-book-api/internal/migration"
	"log"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate.go [up|down]")
	}
	db := config.InitPostgresDB(cfg)
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
