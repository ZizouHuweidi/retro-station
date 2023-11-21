package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/dto/v1"
)

// https://echo.labstack.com/guide/response/
type GetProductsResponseDto struct {
	Products *utils.ListResult[*dtoV1.ProductDto]
}
