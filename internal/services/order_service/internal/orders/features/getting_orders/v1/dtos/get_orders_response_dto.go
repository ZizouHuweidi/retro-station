package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	dtosV1 "github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/orders/dtos/v1"
)

type GetOrdersResponseDto struct {
	Orders *utils.ListResult[*dtosV1.OrderReadDto]
}
