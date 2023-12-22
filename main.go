package main

import (
	"log"
    "os"
    "github.com/rs/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

    config := Config {
        Port: os.Getenv("PORT"),
        CorsOptions: cors.Options{
            AllowedOrigins: []string{os.Getenv("CLIENT_ORIGIN_URL")},
            AllowedMethods: []string{"GET"},
            AllowedHeaders: []string{"Content-Type", "Authorization"},
            MaxAge: 86400,
        },
        Audience: os.Getenv("AUTH0_AUDIENCE"),
        Domain: os.Getenv("AUTH0_DOMAIN"),
    }

    app := App{Config: config}

    app.StartServer();
}
