package middleware

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"log-collection-api/pkg/helpers"

	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
    "github.com/pkg/errors"
)

const (
	missingJWTErrorMessage = "Requires authentication"
	invalidJWTErrorMessage = "Bad credentials"
)

func Logging(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL.Path)
        f(w, r)
    }
}

func ValidateJWT(audience, domain string) func(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://" + domain + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}
    log.Println(issuerURL)

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}


	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)
		if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
			errorMessage := missingJWTErrorMessage//ErrorMessage{Message: missingJWTErrorMessage}
			helpers.WriteJSON(w, http.StatusUnauthorized, errorMessage)
			return
		}
		if errors.Is(err, jwtmiddleware.ErrJWTInvalid) {
			errorMessage := invalidJWTErrorMessage//ErrorMessage{Message: invalidJWTErrorMessage}
			helpers.WriteJSON(w, http.StatusUnauthorized, errorMessage)
			return
		}
        helpers.WriteJSON(w, 200, "success")
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}
