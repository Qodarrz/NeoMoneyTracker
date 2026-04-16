package main

import (
	"log"

	"github.com/Qodarrz/go-gin-air/config"
	"github.com/Qodarrz/go-gin-air/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Connect to database
	config.ConnectDatabase()

	r := router.NewRouter()
	r.Run(":8080")
}
