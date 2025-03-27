package routes

import (
	"github.com/MagnusHLund/VisitorCounter/internal/services"

	"github.com/go-chi/chi/v5"
)

func SetupPageRoutes(app *services.VisitorService) chi.Router {
	router := chi.NewRouter()

	// Use app's handler methods
	router.Post("/", app.CreatePageHandler)

	return router
}
