package middleware

import (
    "strings"
    "context"
    "log"
    "net/http"
    "net/url"
    "os"
    "time"
    jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type CustomClaims struct {
    Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
    return nil
}

func (c CustomClaims) HasScope(expectedScope string) bool {
    result := strings.Split(c.Scope, " ")
    for i := range result {
        if result[i] == expectedScope {
            return true
        }
    }
    return false
}

func EnsureValidToken() func(next http.Handler) http.Handler {
    issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
    if err != nil {
        log.Fatalf("failed to read ussuer url: %v", err)
    }
    provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
    jwtValidator, err := validator.New(
        provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
    )
    if err != nil {
        log.Fatalf("validators failed to setup")
    }
    
	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}
