package wire

import (
	"github.com/MagnusHLund/VisitorCounter/internal/config"
	"github.com/MagnusHLund/VisitorCounter/internal/handlers"
	"github.com/MagnusHLund/VisitorCounter/internal/repositories"
	"github.com/MagnusHLund/VisitorCounter/internal/services"
	"github.com/MagnusHLund/VisitorCounter/internal/utils"
)

type App struct {
	Config       *Config
	Services     *Services
	Handlers     *Handlers
	Repositories *Repositories
	Utils        *Utils
}

type Config struct {
	Config *config.Config
}

type Handlers struct {
	PageHandler *handlers.PageHandler
}

type Services struct {
	PageService    *services.PageService
	VisitorService *services.VisitorService
}

type Repositories struct {
	PageRepository    *repositories.PageRepository
	VisitorRepository *repositories.VisitorRepository
}

type Utils struct {
	RequestUtils *utils.RequestUtils
}

func NewApp(
	config *config.Config,
	pageHandler *handlers.PageHandler,
	pageService *services.PageService,
	visitorService *services.VisitorService,
	pageRepository *repositories.PageRepository,
	visitorRepository *repositories.VisitorRepository,
	requestUtils *utils.RequestUtils,
) *App {
	return &App{
		Config: &Config{
			Config: config,
		},
		Handlers: &Handlers{
			PageHandler: pageHandler,
		},
		Services: &Services{
			PageService:    pageService,
			VisitorService: visitorService,
		},
		Repositories: &Repositories{
			PageRepository:    pageRepository,
			VisitorRepository: visitorRepository,
		},
		Utils: &Utils{
			RequestUtils: requestUtils,
		},
	}
}
