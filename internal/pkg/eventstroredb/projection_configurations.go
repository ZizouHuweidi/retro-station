package eventstroredb

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/es/contracts/projection"
)

type ProjectionsConfigurations struct {
	Projections []projection.IProjection
}
