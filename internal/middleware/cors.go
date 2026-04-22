package middleware

import (
	"net/http"
	"os"
)

// CORS middleware para permitir peticiones desde el frontend
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtener orígenes permitidos de variable de entorno
		allowedOrigin := os.Getenv("ALLOWED_ORIGINS")
		if allowedOrigin == "" {
			allowedOrigin = "*" // Permitir todos en desarrollo
		}

		// Headers CORS
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Manejar preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
