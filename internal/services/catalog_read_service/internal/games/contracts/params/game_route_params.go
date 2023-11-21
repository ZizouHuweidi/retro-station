package params

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"go.uber.org/fx"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/contracts"
)

type GameRouteParams struct {
	fx.In

	CatalogsMetrics *contracts.CatalogsMetrics
	Logger          logger.Logger
	GamesGroup      *echo.Group `name:"game-echo-group"`
	Validator       *validator.Validate
}
