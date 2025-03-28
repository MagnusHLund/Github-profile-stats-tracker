package routes

import (
	"github.com/MagnusHLund/VisitorCounter/internal/middleware"
	"github.com/MagnusHLund/VisitorCounter/internal/wire"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *wire.App) *chi.Mux {
	router := chi.NewRouter()

	applyMiddleware(router)

	pageHandler := handlers.NewPageHandler(app.Handlers.PageHandler.DB)
	visitorHandler := handlers.NewVisitorHandler(app.Handlers.VisitorHandler.DB)

	router.Mount("/page", SetupPageRoutes(pageHandler))
	router.Mount("/visitor", SetupVisitorRoutes(visitorHandler))

	return router
}

func applyMiddleware(router *chi.Mux) {
	router.Use(middleware.ExceptionMiddleware)
	router.Use(middleware.JSONMiddleware)
}
