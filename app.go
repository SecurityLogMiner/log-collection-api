package main 

import (
    "log-collection-api/pkg/routes"
	"log"
	"net/http"

    "github.com/rs/cors"
    "github.com/unrolled/secure"
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

    router := routes.Router(app.Config.Audience, app.Config.Domain)
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

