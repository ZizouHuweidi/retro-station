package params

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/contracts"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type ProductRouteParams struct {
	fx.In

	CatalogsMetrics *contracts.CatalogsMetrics
	Logger          logger.Logger
	ProductsGroup   *echo.Group `name:"product-echo-group"`
	Validator       *validator.Validate
}
