package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/sajin-shrestha/gormTemplate/utils"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Forbidden: empty token", http.StatusForbidden)
			return
		}

		// Remove the 'Bearer ' prefix from the token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Forbidden: invalid token", http.StatusForbidden)
			return
		}

		// Set the claims in the context
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
