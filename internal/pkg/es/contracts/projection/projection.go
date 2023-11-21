package projection

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/pkg/es/models"
)

type IProjection interface {
	ProcessEvent(ctx context.Context, streamEvent *models.StreamEvent) error
}
