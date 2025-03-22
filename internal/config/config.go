package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func GetEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func GetEnvAsInt(key string, fallback int) int {
	valStr := os.Getenv(key)
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return fallback
}

func GetEnvAsBool(key string, fallback bool) bool {
	valStr := os.Getenv(key)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return fallback
}

func GetEnvOrPanic(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Required env var %s is missing", key)
	}
	return val
}
