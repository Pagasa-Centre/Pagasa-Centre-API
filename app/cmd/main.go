package main

import (
	"fmt"
	authentication "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/config"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/router"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
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

	// Create a zap logger.
	// You can choose zap.NewDevelopment() for local development or zap.NewProduction() for production.
	var logger *zap.Logger
	if cfg.Env == "prod" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		fmt.Printf("Failed to create logger: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Printf("Error syncing logger: %v\n", err)
		}
	}() // flushes any buffered log entries

	// Convert the logger to a sugared logger for a more ergonomic API.
	sugaredLogger := logger.Sugar()

	authService := authentication.NewService(*sugaredLogger)

	mux := router.New(*sugaredLogger, authService)

	//log.Printf("Server starting on port %s", cfg.Port)
	sugaredLogger.Infof("Server starting on port %s", cfg.Port)

	addr := fmt.Sprintf(":%s", cfg.Port)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
