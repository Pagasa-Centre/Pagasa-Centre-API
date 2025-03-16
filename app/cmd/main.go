package main

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/config"
	"github.com/joho/godotenv"
	"log"
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

	// Now use cfg.Port, cfg.DatabaseURL, etc. to configure your app.
	log.Printf("Server starting on port %s", cfg.Port)
	// Initialize your database, router, etc.
}
