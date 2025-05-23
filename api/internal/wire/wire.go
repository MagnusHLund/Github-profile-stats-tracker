//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/config"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/database"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/handlers"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/mappers"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/repositories"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/services"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/utils"

	"github.com/google/wire"
)

// CreateApp initializes all dependencies
func CreateApp() (*App, error) {
	wire.Build(
		config.NewConfig,
		database.NewDatabase,
		handlers.NewPageHandler,
		services.NewPageService,
		services.NewVisitorService,
		services.NewHashingService,
		services.NewImageService,
		repositories.NewPageRepository,
		repositories.NewVisitorRepository,
		utils.NewRequestUtils,
		utils.NewSvgUtils,
		utils.NewHelperUtils,
		utils.NewFileUtils,
		mappers.NewQueryParameterMapper,
		NewApp,
	)

	return nil, nil
}
