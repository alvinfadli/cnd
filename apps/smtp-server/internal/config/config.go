package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Domain           string
	Port             string
	AllowInsecureAuth bool
	MaxRecipients    int
}

func LoadConfig() *Config {
	cfg := &Config{
		Domain:        "localhost",
		Port:          "8000",
		MaxRecipients: 50,
	}

	if v := os.Getenv("SMTP_DOMAIN"); v != "" {
		cfg.Domain = v
	}
	if v := os.Getenv("SMTP_PORT"); v != "" {
		cfg.Port = v
	}
	if v := os.Getenv("SMTP_ALLOW_INSECURE_AUTH"); v != "" {
		parsed, err := strconv.ParseBool(v)
		if err != nil {
			log.Printf("Invalid value for SMTP_ALLOW_INSECURE_AUTH=%s, defaulting to false", v)
			cfg.AllowInsecureAuth = false
		} else {
			cfg.AllowInsecureAuth = parsed
		}
	}

	return cfg
}
