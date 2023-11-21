package queries

import "github.com/zizouhuweidi/retro-station/internal/pkg/utils"

// Ref: https://golangbot.com/inheritance/

type GetProducts struct {
	*utils.ListQuery
}

func NewGetProducts(query *utils.ListQuery) *GetProducts {
	return &GetProducts{ListQuery: query}
}
