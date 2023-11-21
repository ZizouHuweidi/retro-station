package consumer

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
)

type Consumer interface {
	Start(ctx context.Context) error
	Stop() error
	ConnectHandler(handler ConsumerHandler)
	IsConsumed(func(message types.IMessage))
	GetName() string
}
