package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN string
	Port  string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system env")
	}

	return Config{
		DBDSN: getEnv(
			"DB_DSN",
			"root:password@tcp(127.0.0.1:3306)/go_wallet?parseTime=true",
		),
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}