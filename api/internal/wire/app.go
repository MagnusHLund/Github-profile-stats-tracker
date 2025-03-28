package wire

import (
	"github.com/MagnusHLund/VisitorCounter/internal/config"
	"github.com/MagnusHLund/VisitorCounter/internal/handlers"
	"github.com/MagnusHLund/VisitorCounter/internal/repositories"
	"github.com/MagnusHLund/VisitorCounter/internal/services"
)

type App struct {
	config       *Config
	Services     *Services
	Handlers     *Handlers
	Repositories *Repositories
}

type Config struct {
	Config *config.Config
}

type Handlers struct {
	PageHandler    *handlers.PageHandler
	VisitorHandler *handlers.VisitorHandler
}

type Services struct {
	PageService    *services.PageService
	VisitorService *services.VisitorService
}

type Repositories struct {
	PageRepository    *repositories.PageRepository
	VisitorRepository *repositories.VisitorRepository
}

func NewApp(
	config *config.Config,
	pageHandler *handlers.PageHandler,
	pageService *services.PageService,
	visitorHandler *handlers.VisitorHandler,
	visitorService *services.VisitorService,
	pageRepository *repositories.PageRepository,
	visitorRepository *repositories.VisitorRepository,
) *App {
	return &App{
		config: &Config{
			Config: config,
		},
		Handlers: &Handlers{
			PageHandler:    pageHandler,
			VisitorHandler: visitorHandler,
		},
		Services: &Services{
			PageService:    pageService,
			VisitorService: visitorService,
		},
		Repositories: &Repositories{
			PageRepository:    pageRepository,
			VisitorRepository: visitorRepository,
		},
	}
}
