package main

import (
	"log"

	"github.com/alvinfadli/cnd/apps/smtp-server/internal/config"
	"github.com/alvinfadli/cnd/apps/smtp-server/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	server.Start(cfg)
}