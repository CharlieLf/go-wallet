package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN  string
	Port   string
	APIKeys []string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system env")
	}

	return Config{
		DBDSN: getEnv("DB_DSN", ""),
		Port:  getEnv("PORT", "8080"),
		APIKeys: strings.Split(
			getEnv("API_KEYS", ""),
			",",
		),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}