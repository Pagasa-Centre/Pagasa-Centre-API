package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"PORT" yaml:"port" validate:"required"`
	DatabaseURL      string `mapstructure:"DATABASE_URL" yaml:"database_url" validate:"required"`
	LogLevel         string `mapstructure:"LOG_LEVEL" yaml:"log_level"`
	Env              string `mapstructure:"ENV" yaml:"env" validate:"required"`
	JwtSecret        string `mapstructure:"JWT_SECRET" yaml:"jwt_secret" validate:"required"`
	YoutubeAPIKey    string `mapstructure:"YOUTUBE_API_KEY" yaml:"youtube_api_key" validate:"required"`
	YoutubeChannelID string `mapstructure:"YOUTUBE_CHANNEL_ID" yaml:"youtube_channel_id" validate:"required"`
}

// LoadConfig loads configuration from the OS environment and, if not in production,
// from a .env file at the root of the repository.
func LoadConfig() (*Config, error) {
	// Check if running in production.
	// When ENV is "prod", we assume all necessary environment variables are set.
	// Otherwise, load variables from the .env file.
	if os.Getenv("ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: no .env file found, relying on OS environment variables: %v", err)
		}
	}

	// Use Viper to read environment variables.
	viper.AutomaticEnv()

	// Set a default value for ENV if it hasn't been set.
	if viper.GetString("ENV") == "" {
		viper.Set("ENV", "dev")
	}

	// Create a Config instance with values from environment variables.
	cfg := Config{
		Port:             viper.GetString("PORT"),
		DatabaseURL:      viper.GetString("DATABASE_URL"),
		LogLevel:         viper.GetString("LOG_LEVEL"),
		JwtSecret:        viper.GetString("JWT_SECRET"),
		YoutubeAPIKey:    viper.GetString("YOUTUBE_API_KEY"),
		YoutubeChannelID: viper.GetString("YOUTUBE_CHANNEL_ID"),
		Env:              viper.GetString("ENV"),
	}

	// Validate the config.
	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &cfg, nil
}
