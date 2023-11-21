package consumer

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
)

type BusControl interface {
	// Start starts all consumers
	Start(ctx context.Context) error
	// Stop stops all consumers
	Stop() error

	IsConsumed(func(message types.IMessage))
}
