package routes

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/middleware"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/wire"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *wire.App) *chi.Mux {
	router := chi.NewRouter()

	applyMiddleware(router)

	pageHandler := app.Handlers.PageHandler

	router.Mount("/page", SetupPageRoutes(pageHandler))

	return router
}

func applyMiddleware(router *chi.Mux) {
	router.Use(middleware.ExceptionMiddleware)
}
