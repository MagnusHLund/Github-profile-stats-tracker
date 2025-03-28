package routes

import (
	"github.com/MagnusHLund/VisitorCounter/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupPageRoutes(app *handlers.PageHandler) chi.Router {
	router := chi.NewRouter()

	// Use app's handler methods
	router.Post("/", app.CreatePage)

	return router
}
