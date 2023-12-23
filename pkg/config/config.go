package config

import (
    "os"
    "github.com/rs/cors"
)

func CorsOptions() cors.Options {
    return cors.Options {
        AllowedOrigins: []string{os.Getenv("CLIENT_ORIGIN_URL")},
        AllowedMethods: []string{"GET"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
        MaxAge: 86400,
    }
}
