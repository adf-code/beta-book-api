package main

import (
	"beta-book-api/config"
	api "beta-book-api/internal/delivery/http"
	"beta-book-api/internal/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	db := config.InitPostgresDB(cfg)
	repo := repository.NewBookRepo(db)
	handler := api.NewBookHandler(repo)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Println("🟢 Server started on :8080 | Press Ctrl+C to stop")
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Println("❌ Server stopped unexpectedly:", err)
	}
}
