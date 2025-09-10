package main

import (
	"log"

	"github.com/alvinfadli/cnd/apps/smtp-server/internal/config"
	"github.com/alvinfadli/cnd/apps/smtp-server/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	log.Printf("Loaded config: Domain=%s, Port=%s, AllowInsecureAuth=%v",
		cfg.Domain, cfg.Port, cfg.AllowInsecureAuth,
	)

	server.Start(cfg)
}
