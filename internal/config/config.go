package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config holds all configuration for your application.
type Config struct {
	Port string `mapstructure:"port"`
	//DatabaseURL string `mapstructure:"database_url"`
	LogLevel string `mapstructure:"log_level"`
}

// LoadConfig reads configuration from file and environment variables.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./internal/config")

	// Set defaults
	viper.SetDefault("port", "8080")
	//viper.SetDefault("database_url", "postgres://user:password@localhost:5432/dbname?sslmode=disable")
	viper.SetDefault("log_level", "info")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &cfg, nil
}
