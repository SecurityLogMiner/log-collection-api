package config

import (
    "os"
    "github.com/rs/cors"
)

func CorsOptions() cors.Options {
    return cors.Options {
        //AllowedCredentials: []string{"Access-Control-Allow-Credentials", "true"},
        AllowedOrigins: []string{os.Getenv("CLIENT_ORIGIN_URL")},
        AllowedMethods: []string{"GET"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
        MaxAge: 86400,
    }
}
