package consumer

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
)

type ConsumerHandler interface {
	Handle(ctx context.Context, consumeContext types.MessageConsumeContext) error
}
