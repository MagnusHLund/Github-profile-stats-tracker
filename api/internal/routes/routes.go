package routes

import (
	"github.com/MagnusHLund/VisitorCounter/internal/services"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *services.VisitorService) *chi.Mux {
	router := chi.NewRouter()

	router.Use(loggingMiddleware)

	router.Mount("/page", SetupPageRoutes(app))
	router.Mount("/visitor", SetupVisitorRoutes(app))

	return router
}

// TODO: Put middleware in its own package
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details
		log.Printf("Method: %s, Path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
