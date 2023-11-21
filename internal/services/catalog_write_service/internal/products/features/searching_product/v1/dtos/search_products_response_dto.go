package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/dto/v1"
)

type SearchProductsResponseDto struct {
	Products *utils.ListResult[*dtoV1.ProductDto]
}
