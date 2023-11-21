package test

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/test"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"

	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/shared/app"
	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/shared/configurations/orders"
)

type OrdersTestApplication struct {
	*app.OrdersApplication
	tb fxtest.TB
}

func NewOrdersTestApplication(
	tb fxtest.TB,
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	environment environemnt.Environment,
) *OrdersTestApplication {
	testApp := test.NewTestApplication(
		tb,
		providers,
		decorates,
		options,
		logger,
		environment,
	)

	orderApplication := &app.OrdersApplication{
		OrdersServiceConfigurator: orders.NewOrdersServiceConfigurator(testApp),
	}

	return &OrdersTestApplication{
		OrdersApplication: orderApplication,
		tb:                tb,
	}
}
