package integrationEvents

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	dtosV1 "github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/orders/dtos/v1"

	uuid "github.com/satori/go.uuid"
)

type OrderCreatedV1 struct {
	*types.Message
	*dtosV1.OrderReadDto
}

func NewOrderCreatedV1(orderReadDto *dtosV1.OrderReadDto) *OrderCreatedV1 {
	return &OrderCreatedV1{
		OrderReadDto: orderReadDto,
		Message:      types.NewMessage(uuid.NewV4().String()),
	}
}
