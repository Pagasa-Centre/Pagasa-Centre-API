package config

import (
	"fmt"

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

// LoadConfig reads configuration from file and environment variables.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./internal/config")

	// Set defaults
	viper.SetDefault("port", "8080")
	viper.SetDefault("log_level", "info")
	viper.SetDefault("env", "dev")

	// Bind environment variables explicitly.
	// This ensures that values from the OS env (like those loaded by godotenv) are used.
	err := viper.BindEnv("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	err = viper.BindEnv("JWT_SECRET")
	if err != nil {
		return nil, err
	}

	err = viper.BindEnv("YOUTUBE_API_KEY")
	if err != nil {
		return nil, err
	}

	err = viper.BindEnv("YOUTUBE_CHANNEL_ID")
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	if err := validate[Config](cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &cfg, nil
}

// validate validates the config against the struct tags.
func validate[T any](target T) error {
	return validator.New().Struct(target)
}
