package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

func ExceptionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer handleException(w, r)
		next.ServeHTTP(w, r)
	})
}

func handleException(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		logError(err, r.Method, r.URL.String())
		sendResponse(w)
	}
}

// TODO: Add Timestamp to log, add log files
func logError(err interface{}, httpMethod string, url string) {
	log.Printf("Panic recovered: %v | Method: %s | URL: %s", err, httpMethod, url)
}

// TODO: add better responses
func sendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   "Internal Server Error",
		"message": "An unexpected error occurred",
	})
}
