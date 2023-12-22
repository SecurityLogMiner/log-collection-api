package routes 

import (
	"log-collection-api/middleware"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/api/messages/public", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("hello from %s",r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello from a public endpoint! You don't need to be authenticated to see this."}`))
	})

	// This route is only accessible if the user has a valid access_token.
	router.HandleFunc("/api/messages/private", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
    })

	// This route is only accessible if the user has a
	// valid access_token with the read:messages scope.
	router.HandleFunc("/api/messages/private-scoped", func(w http.ResponseWriter, r *http.Request) {
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
	})
    
    return router
}
