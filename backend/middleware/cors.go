package middleware

import "net/http"

// CORS Middleware to handle cross-origin requests
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin (for development)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Specify allowed headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Specify allowed methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}
