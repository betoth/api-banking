package main

import (
	"api-banking/app"
	"api-banking/config"
	"api-banking/logger"
	"log"
)

func main() {

	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Perform sanity check on the loaded configuration
	if err := config.SanityCheck(cfg); err != nil {
		log.Fatalf("Sanity check failed: %v", err)
	}

	logger.Info("Starting the application")
	app.Start(cfg)
}
