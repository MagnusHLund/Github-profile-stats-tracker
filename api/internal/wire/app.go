package wire

import (
	"github.com/MagnusHLund/VisitorCounter/internal/handlers"
	"github.com/MagnusHLund/VisitorCounter/internal/services"
)

type App struct {
	Services *Services
	Handlers *Handlers
}

type Handlers struct {
	PageHandler    *handlers.PageHandler
	VisitorHandler *handlers.VisitorHandler
}

type Services struct {
	PageService    *services.PageService
	VisitorService *services.VisitorService
}

func NewApp(
	pageHandler *handlers.PageHandler,
	visitorHandler *handlers.VisitorHandler,
	pageService *services.PageService,
	visitorService *services.VisitorService,
) *App {
	return &App{
		Handlers: &Handlers{
			PageHandler:    pageHandler,
			VisitorHandler: visitorHandler,
		},
		Services: &Services{
			PageService:    pageService,
			VisitorService: visitorService,
		},
	}
}
