package app

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/shared/configurations/orders"
	"go.uber.org/fx"
)

type OrdersApplication struct {
	*orders.OrdersServiceConfigurator
}

func NewOrdersApplication(
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	environment environemnt.Environment,
) *OrdersApplication {
	app := fxapp.NewApplication(providers, decorates, options, logger, environment)
	return &OrdersApplication{
		OrdersServiceConfigurator: orders.NewOrdersServiceConfigurator(app),
	}
}
