package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in environment or .env file")
	}

	// Initialize migrate
	// path relative to the root of the project where you run the command
	m, err := migrate.New(
		"file://db/migrations",
		dbURL,
	)
	if err != nil {
		log.Fatal("Failed to initialize migrate: ", err)
	}

	// Simple CLI arguments handling
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/migrate/main.go [up|down|force <version>]")
		return
	}

	command := os.Args[1]
	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration Up failed: ", err)
		}
		fmt.Println("Migration UP completed successfully!")
	case "down":
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration Down failed: ", err)
		}
		fmt.Println("Migration DOWN completed successfully!")
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Force command requires a version number")
		}
		version := 0
		fmt.Sscanf(os.Args[2], "%d", &version)
		if err := m.Force(version); err != nil {
			log.Fatal("Migration Force failed: ", err)
		}
		fmt.Println("Migration forced to version ", version)
	default:
		fmt.Println("Unknown command. Use: up, down, or force")
	}
}
