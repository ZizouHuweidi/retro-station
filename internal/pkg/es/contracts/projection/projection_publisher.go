package projection

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/pkg/es/models"
)

type IProjectionPublisher interface {
	Publish(ctx context.Context, streamEvent *models.StreamEvent) error
}
