package config

import (
	"api-banking/infra/env"
	"errors"
	"log"
	"os"
)

// Config holds the configuration values for the application
type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
	ServerAddress string
	ServerPort    string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	err := env.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBSSLMode:     os.Getenv("DB_SSL_MODE"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		ServerPort:    os.Getenv("SERVER_PORT"),
	}, nil
}

// SanityCheck ensures all necessary configuration values are present
func SanityCheck(cfg *Config) error {
	if cfg.DBHost == "" {
		return errors.New("missing DB_HOST")
	}
	if cfg.DBPort == "" {
		return errors.New("missing DB_PORT")
	}
	if cfg.DBUser == "" {
		return errors.New("missing DB_USER")
	}
	if cfg.DBPassword == "" {
		return errors.New("missing DB_PASSWORD")
	}
	if cfg.DBName == "" {
		return errors.New("missing DB_NAME")
	}
	if cfg.DBSSLMode == "" {
		return errors.New("missing DB_SSL_MODE")
	}
	if cfg.ServerAddress == "" {
		return errors.New("missing SERVER_ADDRESS")
	}
	if cfg.ServerPort == "" {
		return errors.New("missing SERVER_PORT")
	}
	return nil
}
