//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/MagnusHLund/VisitorCounter/internal/config"
	"github.com/MagnusHLund/VisitorCounter/internal/database"
	"github.com/MagnusHLund/VisitorCounter/internal/services"

	"github.com/google/wire"
)

// CreateApp initializes all dependencies
func CreateApp() (*services.VisitorService, error) {
	wire.Build(config.NewConfig, database.NewDatabase, services.NewVisitorService)
	return nil, nil
}

// TODO: Add more wire.Build calls for other services and repositories
