package main

import (
	"fmt"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/config"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// Load environment variables from .env file (only in development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	// Now load your app configuration using Viper (which automatically reads env vars)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	mux := router.New()

	log.Printf("Server starting on port %s", cfg.Port)

	addr := fmt.Sprintf(":%s", cfg.Port)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
