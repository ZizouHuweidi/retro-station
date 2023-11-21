package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/products/dto"
)

type SearchProductsResponseDto struct {
	Products *utils.ListResult[*dto.ProductDto]
}
