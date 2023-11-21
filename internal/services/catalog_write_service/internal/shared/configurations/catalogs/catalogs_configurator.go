package catalogs

import (
	"fmt"
	"net/http"

	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	customEcho "github.com/zizouhuweidi/retro-station/internal/pkg/http/custom_echo"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/config"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/configurations"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/configurations/catalogs/infrastructure"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CatalogsServiceConfigurator struct {
	contracts.Application
	infrastructureConfigurator *infrastructure.InfrastructureConfigurator
	productsModuleConfigurator *configurations.ProductsModuleConfigurator
}

func NewCatalogsServiceConfigurator(app contracts.Application) *CatalogsServiceConfigurator {
	infraConfigurator := infrastructure.NewInfrastructureConfigurator(app)
	productModuleConfigurator := configurations.NewProductsModuleConfigurator(app)

	return &CatalogsServiceConfigurator{
		Application:                app,
		infrastructureConfigurator: infraConfigurator,
		productsModuleConfigurator: productModuleConfigurator,
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
	// Product module
	ic.productsModuleConfigurator.ConfigureProductsModule()
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
	// Products CatalogsServiceModule endpoints
	ic.productsModuleConfigurator.MapProductsEndpoints()
}