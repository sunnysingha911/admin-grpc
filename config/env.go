package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	JWTSecret  string
	GRPCUSER   string
)

func LoadEnv() error {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		// Not a fatal error; you might want to proceed if env vars are set elsewhere
	}

	DBHost = getEnv("DB_HOST", "localhost")
	DBUser = getEnv("DB_USER", "user")
	DBPassword = getEnv("DB_PASSWORD", "password")
	DBName = getEnv("DB_NAME", "users_db")
	DBPort = getEnv("DB_PORT", "5432")
	JWTSecret = getEnv("JWT_SECRET", "my-super-secret-key")
	GRPCUSER = getEnv("GRPC_USER", "localhost:50051")

	// Optional: Validate critical env vars
	if DBUser == "" || DBPassword == "" || DBName == "" {
		return fmt.Errorf("critical environment variables missing")
	}

	return nil
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
