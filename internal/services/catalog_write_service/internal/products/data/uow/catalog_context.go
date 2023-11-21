package uow

import (
	data2 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/contracts/data"
)

type catalogContext struct {
	productRepository data2.ProductRepository
}

func (c *catalogContext) Products() data2.ProductRepository {
	return c.productRepository
}
