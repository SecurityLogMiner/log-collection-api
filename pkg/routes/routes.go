package routes 

import (
	"log-collection-api/pkg/middleware"
	"net/http"
	"os"
	"fmt"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func Router(audience, domain string) http.Handler {
	router := http.NewServeMux()

	router.Handle("/api/public", 
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`{"message":"reached a public endpoint."}`))
	}))

	router.Handle("/api/private", middleware.EnsureValidToken()(
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Credentials", "true")
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`{"message":"reached private endpoint"}`))
        }),
    ))

	router.Handle("/api/private-scoped", middleware.EnsureValidToken()(
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

    return router
}

func userHasCertificate(userID string) bool {
	//template
	return false 
}

//basic implementation. will continue to work on it
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")

	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	if !userHasCertificate(userID) {
		http.Error(w, "No certificate found for the user", http.StatusNotFound)
		return
	}

	certificatePath := fmt.Sprintf("/path/to/certificates/%s-certificate.pdf", userID)

	file, err := os.Open(certificatePath)
	if err != nil {
		http.Error(w, "Error opening the certificate file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-certificate.pdf", userID))
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeContent(w, r, certificatePath, file.Stat().ModTime(), file)
}

