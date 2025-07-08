package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBURL    string
	LogLevel string
}

func LoadConfig() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("env not found")
	}
	return &Config{
		Port:     getEnv("PORT", "8080"),
		DBURL:    getEnv("DB_URL", ""),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultValue
}
