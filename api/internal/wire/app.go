package wire

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/config"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/handlers"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/mappers"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/repositories"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/services"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/utils"
)

type App struct {
	Config       *Config
	Services     *Services
	Handlers     *Handlers
	Repositories *Repositories
	Utils        *Utils
	Mappers      *Mappers
}

type Config struct {
	Config *config.Config
}

type Handlers struct {
	PageHandler *handlers.PageHandler
}

type Services struct {
	PageService    *services.PageService
	ImageService   *services.ImageService
	VisitorService *services.VisitorService
	HashingService *services.HashingService
}

type Repositories struct {
	PageRepository    *repositories.PageRepository
	VisitorRepository *repositories.VisitorRepository
}

type Utils struct {
	RequestUtils *utils.RequestUtils
	SvgUtils     *utils.SVGUtils
	HelperUtils  *utils.HelperUtils
	FileUtils    *utils.FileUtils
}

type Mappers struct {
	QueryParameterMapper *mappers.QueryParameterMapper
}

func NewApp(
	config *config.Config,
	pageHandler *handlers.PageHandler,
	pageService *services.PageService,
	visitorService *services.VisitorService,
	imageService *services.ImageService,
	hashingService *services.HashingService,
	pageRepository *repositories.PageRepository,
	visitorRepository *repositories.VisitorRepository,
	requestUtils *utils.RequestUtils,
	svgUtils *utils.SVGUtils,
	helperUtils *utils.HelperUtils,
	fileUtils *utils.FileUtils,
	queryParameterMapper *mappers.QueryParameterMapper,

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
			ImageService:   imageService,
			VisitorService: visitorService,
			HashingService: hashingService,
		},
		Repositories: &Repositories{
			PageRepository:    pageRepository,
			VisitorRepository: visitorRepository,
		},
		Utils: &Utils{
			RequestUtils: requestUtils,
			SvgUtils:     svgUtils,
			HelperUtils:  helperUtils,
			FileUtils:    fileUtils,
		},
		Mappers: &Mappers{
			QueryParameterMapper: queryParameterMapper,
		},
	}
}
