package producer

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/pkg/core/metadata"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
)

type Producer interface {
	PublishMessage(ctx context.Context, message types.IMessage, meta metadata.Metadata) error
	PublishMessageWithTopicName(
		ctx context.Context,
		message types.IMessage,
		meta metadata.Metadata,
		topicOrExchangeName string,
	) error
	IsProduced(func(message types.IMessage))
}
