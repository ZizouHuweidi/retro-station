package params

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/es/contracts/projection"

	"go.uber.org/fx"
)

type OrderProjectionParams struct {
	fx.In

	Projections []projection.IProjection `group:"projections"`
}
