package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken  string
	AppID          string
	MarketHashName string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		TelegramToken:  getEnv("TELEGRAM_TOKEN", "required"),
		AppID:          getEnv("APP_ID", "required"),
		MarketHashName: getEnv("MARKET_HASH_NAME", "required"),
	}
}

func getEnv(key string, required string) string {
	value := os.Getenv(key)
	if required == "required" && value == "" {
		log.Fatalf("Environment variable %s is required", key)
	}
	return value
}