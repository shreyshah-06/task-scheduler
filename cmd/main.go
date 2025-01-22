package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shreyshah-06/taskScheduler/internal/db"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Database configuration
	cfg := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     mustConvertToInt(os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	database := db.InitDB(cfg)
	defer database.Close()

	// Test database connection
	err = database.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Database connection successful!")
}

// Helper function to convert string to int
func mustConvertToInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Invalid integer value: %s", value)
	}
	return val
}
