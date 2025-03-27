package routes

import (
	"github.com/MagnusHLund/VisitorCounter/internal/services"

	"github.com/go-chi/chi/v5"
)

func SetupVisitorRoutes(app *services.VisitorService) chi.Router {
	router := chi.NewRouter()

	// Use app's handler methods
	router.Get("/{page}", app.GetVisitorsHandler)
	router.Post("/", app.CreateVisitorHandler)

	return router
}
