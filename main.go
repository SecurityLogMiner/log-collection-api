package main

import (
	"log"
    "os"
	"github.com/joho/godotenv"
    "log-collection-api/pkg/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}
    config := Config {
        Port: os.Getenv("PORT"),
        CorsOptions: config.CorsOptions(),
        Audience: os.Getenv("AUTH0_AUDIENCE"),
        Domain: os.Getenv("AUTH0_DOMAIN"),
    }
    app := App{Config: config}
    app.StartServer();
}
