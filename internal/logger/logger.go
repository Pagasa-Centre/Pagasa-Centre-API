package logger

import (
	"fmt"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/config"
	"go.uber.org/zap"
	"os"
)

func New(cfg *config.Config) *zap.SugaredLogger {
	var logger *zap.Logger
	var err error

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

	return sugaredLogger
}
