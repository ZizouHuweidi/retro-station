package dtos

import "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/products/dto"

type GetProductByIdResponseDto struct {
	Product *dto.ProductDto `json:"product"`
}
