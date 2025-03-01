package main

import (
	"log"
	"os"

	"github.com/Seew0/Heal-D/internal/server"
	"github.com/Seew0/Heal-D/internal/wire"
	"github.com/joho/godotenv"
)

func init() {
	log.Println("Loading .env file")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router, err := wire.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}


	server := server.NewServer(router)
	server.Start(port)
}
