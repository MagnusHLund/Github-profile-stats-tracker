package routes

import (
	"github.com/MagnusHLund/VisitorCounter/internal/middleware"
	"github.com/MagnusHLund/VisitorCounter/internal/wire"

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
	router.Use(middleware.JSONMiddleware)
}
