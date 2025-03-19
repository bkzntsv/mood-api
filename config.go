package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config holds application configuration loaded from environment or config file
type Config struct {
	Port          string `mapstructure:"PORT" default:"8080"`          // Server port to listen on
	ModelEndpoint string `mapstructure:"MODEL_ENDPOINT" required:"true"` // Required ML model endpoint
}

// LoadConfig loads configuration from the specified path
func LoadConfig(path string) (Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Set default values
	viper.SetDefault("PORT", "8080")

	// Enable reading from environment variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found is OK since we can use env vars
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("unable to decode config: %w", err)
	}

	// Validate required fields
	if config.ModelEndpoint == "" {
		return config, fmt.Errorf("MODEL_ENDPOINT is required")
	}

	return config, nil
}

// GetConfig loads and returns the application configuration
func GetConfig() Config {
	config, err := LoadConfig(".")
	if err != nil {
		log.Printf("Configuration error: %v", err)
		os.Exit(1)
	}

	log.Printf("Configuration loaded successfully. Using port: %s", config.Port)
	return config
}