package middlewares

import (
	"net/http"
	"os"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedToken := os.Getenv("TOKEN")

		requestToken := r.Header.Get("token")
		if requestToken == "" {
			http.Error(w, "Invalid Header", http.StatusUnauthorized)
			return
		}

		if requestToken != expectedToken {
			http.Error(w, "Invalid Header", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
