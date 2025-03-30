//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/MagnusHLund/VisitorCounter/internal/config"
	"github.com/MagnusHLund/VisitorCounter/internal/database"
	"github.com/MagnusHLund/VisitorCounter/internal/handlers"
	"github.com/MagnusHLund/VisitorCounter/internal/repositories"
	"github.com/MagnusHLund/VisitorCounter/internal/services"
	"github.com/MagnusHLund/VisitorCounter/internal/utils"

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
		repositories.NewPageRepository,
		repositories.NewVisitorRepository,
		utils.NewRequestUtils,
		NewApp,
	)

	return nil, nil
}
