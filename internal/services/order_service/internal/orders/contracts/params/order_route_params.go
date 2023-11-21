package params

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/shared/contracts"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type OrderRouteParams struct {
	fx.In

	OrdersMetrics *contracts.OrdersMetrics
	Logger        logger.Logger
	OrdersGroup   *echo.Group `name:"order-echo-group"`
	Validator     *validator.Validate
}
