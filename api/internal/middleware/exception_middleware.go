package middleware

import (
	"encoding/json"
	"net/http"
)

func ExceptionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer handleException(w)
		next.ServeHTTP(w, r)
	})
}

// TODO: Add logging and better response
func handleException(w http.ResponseWriter) {
	if err := recover(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Internal Server Error",
		})
	}
}
