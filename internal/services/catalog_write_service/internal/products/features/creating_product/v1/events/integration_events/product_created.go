package integration_events

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/dto/v1"

	uuid "github.com/satori/go.uuid"
)

type ProductCreatedV1 struct {
	*types.Message
	*dtoV1.ProductDto
}

func NewProductCreatedV1(productDto *dtoV1.ProductDto) *ProductCreatedV1 {
	return &ProductCreatedV1{
		ProductDto: productDto,
		Message:    types.NewMessage(uuid.NewV4().String()),
	}
}
