package routes 

import (
	"log-collection-api/pkg/middleware"
	"net/http"

	//jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	//"github.com/auth0/go-jwt-middleware/v2/validator"
)

func Router(audience, domain string) http.Handler {
	router := http.NewServeMux()

<<<<<<< HEAD
	router.HandleFunc("/api/public", func(w http.ResponseWriter, r *http.Request) {
=======
    router.Handle("/api/private", 
        middleware.ValidateJWT(audience, domain)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            log.Println(r)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`{"message":"reached private endpoint"}`))
        })))

    /*
	router.HandleFunc("/api/messages/public", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("hello from %s",r)
>>>>>>> parent of 6193b6e (jwt validated. scope test passed)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello from a public endpoint!"}`))
	})

	// This route is only accessible if the user has a valid access_token.
<<<<<<< HEAD
	router.Handle("/api/private", middleware.EnsureValidToken()(
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Credentials", "true")
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`{"message":"reached private endpoint"}`))
        }),
    ))


	// This route is only accessible if the user has a
	// valid access_token with the read:messages scope.
	router.Handle("/api/private-scoped",middleware.EnsureValidToken()(
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Content-Type", "application/json")

			token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

			claims := token.CustomClaims.(*middleware.CustomClaims)
			if !claims.HasScope("read:messages") {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"message":"Insufficient scope."}`))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Hello from a private, scoped, endpoint!"}`))
	    }),
    ))

=======
	router.HandleFunc("/api/messages/private", middleware.Logging(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message":"reached private endpoint"}`))
    }))

	// This route is only accessible if the user has a
	// valid access_token with the read:messages scope.
	router.HandleFunc("/api/messages/private-scoped", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        log.Println(r.Context())
        token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
        claims := token.CustomClaims.(*middleware.CustomClaims)
        log.Println(claims)
        if !claims.Permissions {
            w.WriteHeader(http.StatusForbidden)
            w.Write([]byte(`{"message":"Insufficient scope."}`))
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
	})
    */
    
>>>>>>> parent of 6193b6e (jwt validated. scope test passed)
    return router
}

