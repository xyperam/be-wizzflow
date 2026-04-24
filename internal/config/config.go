package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	JWTSecret  string
}

func LoadConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := &Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "5433"), // Sesuai jalur evakuasi kemarin
		DBName:     getEnv("DB_NAME", "wizzflow"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
