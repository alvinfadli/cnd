package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Domain            string
	Port              string
	Username          string
	Password          string
	MaxMessageBytes   int
	MaxRecipients     int
	WriteTimeout      time.Duration
	ReadTimeout       time.Duration
	AllowInsecureAuth bool
}

func Load() (*Config, error) {
	config := &Config{
		Domain:            getEnv("SMTP_DOMAIN", "example.com"),
		Port:              getEnv("SMTP_PORT", "2525"),
		Username:          getEnv("SMTP_USERNAME", ""),
		Password:          getEnv("SMTP_PASSWORD", ""),
		MaxMessageBytes:   getEnvInt("SMTP_MAX_MESSAGE_BYTES", 1024*1024),
		MaxRecipients:     getEnvInt("SMTP_MAX_RECIPIENTS", 50),
		WriteTimeout:      time.Duration(getEnvInt("SMTP_WRITE_TIMEOUT", 10)) * time.Second,
		ReadTimeout:       time.Duration(getEnvInt("SMTP_READ_TIMEOUT", 10)) * time.Second,
		AllowInsecureAuth: getEnvBool("SMTP_ALLOW_INSECURE_AUTH", true),
	}

	if config.Username == "" {
		return nil, errors.New("SMTP_USERNAME environment variable is required")
	}
	if config.Password == "" {
		return nil, errors.New("SMTP_PASSWORD environment variable is required")
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}