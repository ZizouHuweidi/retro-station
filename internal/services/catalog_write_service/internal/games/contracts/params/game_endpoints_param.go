package params

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/web/route"
	"go.uber.org/fx"
)

type GamesEndpointsParams struct {
	fx.In

	Endpoints []route.Endpoint `group:"game-routes"`
}
