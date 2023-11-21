package params

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/web/route"

	"go.uber.org/fx"
)

type ProductsEndpointsParams struct {
	fx.In

	Endpoints []route.Endpoint `group:"product-routes"`
}
