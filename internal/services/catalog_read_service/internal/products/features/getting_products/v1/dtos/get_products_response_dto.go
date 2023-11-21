package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/products/dto"
)

type GetProductsResponseDto struct {
	Products *utils.ListResult[*dto.ProductDto]
}
