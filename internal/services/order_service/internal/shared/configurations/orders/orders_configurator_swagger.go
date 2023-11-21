package orders

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	customEcho "github.com/zizouhuweidi/retro-station/internal/pkg/http/custom_echo"

	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/docs"
)

func (ic *OrdersServiceConfigurator) configSwagger(routeBuilder *customEcho.RouteBuilder) {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Orders Service Api"
	docs.SwaggerInfo.Description = "Orders Service Api."

	routeBuilder.RegisterRoutes(func(e *echo.Echo) {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	})
}
