package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"siap_app/internal/app/helpers"
	"strings"
)

// AuthorizationMiddleware checks the authorization header and verifies the token.
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromHeader(r)
		if tokenString == "" {
			writeUnauthorizedResponse(w)
			return
		}

		tokenData, err := helpers.VerifyToken(tokenString)
		if err != nil {
			writeUnauthorizedResponse(w)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tokenData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			return parts[1]
		}
	}
	return ""
}

func writeUnauthorizedResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "Unauthorized",
		"status":     "Failed",
		"statusCode": http.StatusUnauthorized,
	})
}
