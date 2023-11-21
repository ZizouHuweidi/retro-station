package integration_events

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	dto "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/dto/v1"

	uuid "github.com/satori/go.uuid"
)

type ProductUpdatedV1 struct {
	*types.Message
	*dto.ProductDto
}

func NewProductUpdatedV1(productDto *dto.ProductDto) *ProductUpdatedV1 {
	return &ProductUpdatedV1{
		Message:    types.NewMessage(uuid.NewV4().String()),
		ProductDto: productDto,
	}
}
