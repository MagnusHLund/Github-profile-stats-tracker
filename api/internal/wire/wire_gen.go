// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/config"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/database"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/handlers"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/mappers"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/repositories"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/services"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/utils"
)

// Injectors from wire.go:

// CreateApp initializes all dependencies
func CreateApp() (*App, error) {
	configConfig := config.NewConfig()
	db, err := database.NewDatabase(configConfig)
	if err != nil {
		return nil, err
	}
	pageRepository := repositories.NewPageRepository(db)
	visitorRepository := repositories.NewVisitorRepository(db)
	hashingService := services.NewHashingService()
	visitorService := services.NewVisitorService(visitorRepository, hashingService)
	fileUtils := utils.NewFileUtils()
	svgUtils := utils.NewSvgUtils(fileUtils)
	imageService := services.NewImageService(svgUtils)
	pageService := services.NewPageService(pageRepository, visitorService, imageService)
	requestUtils := utils.NewRequestUtils()
	helperUtils := utils.NewHelperUtils()
	queryParameterMapper := mappers.NewQueryParameterMapper(helperUtils)
	pageHandler := handlers.NewPageHandler(pageService, requestUtils, visitorService, queryParameterMapper)
	app := NewApp(configConfig, pageHandler, pageService, visitorService, imageService, hashingService, pageRepository, visitorRepository, requestUtils, svgUtils, helperUtils, fileUtils, queryParameterMapper)
	return app, nil
}
