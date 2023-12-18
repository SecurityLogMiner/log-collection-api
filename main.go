package main

import (
	"01-Authorization-RS256/middleware"
	"log"
	"net/http"
    "os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
    "github.com/rs/cors"
    "github.com/unrolled/secure"
	"github.com/joho/godotenv"
)

type Config struct {
    Port string
    CorsOptions cors.Options
    Audience string
    Domain string
}

type App struct {
    Config Config
}

func (app *App) StartServer() {
	router := http.NewServeMux()
	// This route is always accessible.
	router.Handle("/api/messages/public", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("hello from %s",r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello from a public endpoint! You don't need to be authenticated to see this."}`))
	}))

	// This route is only accessible if the user has a valid access_token.
	router.Handle("/api/messages/private", middleware.EnsureValidToken()(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
		}),
	))

	// This route is only accessible if the user has a
	// valid access_token with the read:messages scope.
	router.Handle("/api/messages/private-scoped", middleware.EnsureValidToken()(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

			claims := token.CustomClaims.(*middleware.CustomClaims)
			if !claims.HasScope("read:messages") {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"message":"Insufficient scope."}`))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
		}),
	))

    corsMiddleware := cors.New(app.Config.CorsOptions)
    routerWithCORS := corsMiddleware.Handler(router)

    secureMiddleware := secure.New()//app.Config.SecureOptions)
    finalHandler := secureMiddleware.Handler(routerWithCORS)
    
    server := &http.Server{
        Addr: ":" + app.Config.Port,
        Handler: finalHandler,
    }

    log.Printf("API server running on %s", server.Addr)
    log.Fatal(server.ListenAndServe())
}

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
    /*
	log.Print("Server listening on http://localhost:6060")
	if err := http.ListenAndServe("0.0.0.0:6060", router); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
    */
    app.StartServer();
}
