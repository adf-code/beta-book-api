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
	deliveryHttp "beta-book-api/internal/delivery/http"
	"beta-book-api/internal/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	db := config.InitPostgresDB(cfg)
	repo := repository.NewBookRepo(db)
	handler := deliveryHttp.SetupHandler(repo)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Println("🟢 Server started on :8080 | Press Ctrl+C to stop")
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Println("❌ Server stopped unexpectedly:", err)
	}
}
