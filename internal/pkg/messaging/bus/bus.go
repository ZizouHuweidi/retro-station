package bus

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/consumer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/producer"
)

type Bus interface {
	producer.Producer
	consumer.BusControl
	consumer.ConsumerConnector
}
