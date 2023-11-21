package catalogs

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	customEcho "github.com/zizouhuweidi/retro-station/internal/pkg/http/custom_echo"
	"gorm.io/gorm"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/config"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/configurations"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/configurations/catalogs/infrastructure"
)

type CatalogsServiceConfigurator struct {
	contracts.Application
	infrastructureConfigurator *infrastructure.InfrastructureConfigurator
	gamesModuleConfigurator    *configurations.GamesModuleConfigurator
}

func NewCatalogsServiceConfigurator(app contracts.Application) *CatalogsServiceConfigurator {
	infraConfigurator := infrastructure.NewInfrastructureConfigurator(app)
	gameModuleConfigurator := configurations.NewGamesModuleConfigurator(app)

	return &CatalogsServiceConfigurator{
		Application:                app,
		infrastructureConfigurator: infraConfigurator,
		gamesModuleConfigurator:    gameModuleConfigurator,
	}
}

func (ic *CatalogsServiceConfigurator) ConfigureCatalogs() {
	// Shared
	// Infrastructure
	ic.infrastructureConfigurator.ConfigInfrastructures()

	// Shared
	// Catalogs configurations
	ic.ResolveFunc(func(gorm *gorm.DB) error {
		err := ic.migrateCatalogs(gorm)
		if err != nil {
			return err
		}

		return nil
	})

	// Modules
	// Game module
	ic.gamesModuleConfigurator.ConfigureGamesModule()
}

func (ic *CatalogsServiceConfigurator) MapCatalogsEndpoints() {
	// Shared
	ic.ResolveFunc(
		func(catalogsServer customEcho.EchoHttpServer, options *config.AppOptions) error {
			catalogsServer.SetupDefaultMiddlewares()

			// Config catalogs root endpoint
			catalogsServer.RouteBuilder().
				RegisterRoutes(func(e *echo.Echo) {
					e.GET("", func(ec echo.Context) error {
						return ec.String(
							http.StatusOK,
							fmt.Sprintf("%s is running...", options.GetMicroserviceNameUpper()),
						)
					})
				})

			// Config catalogs swagger
			ic.configSwagger(catalogsServer.RouteBuilder())

			return nil
		},
	)

	// Modules
	// Games CatalogsServiceModule endpoints
	ic.gamesModuleConfigurator.MapGamesEndpoints()
}
