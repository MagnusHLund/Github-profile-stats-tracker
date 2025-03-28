package middleware

import (
	"encoding/json"
	"net/http"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if requiresJSON(r.Method) && !isJSONContentType(r) {
			sendUnsupportedMediaTypeResponse(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func requiresJSON(method string) bool {
	return method == http.MethodPost || method == http.MethodPut
}

func isJSONContentType(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

func sendUnsupportedMediaTypeResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnsupportedMediaType)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Unsupported Media Type",
	})
}
