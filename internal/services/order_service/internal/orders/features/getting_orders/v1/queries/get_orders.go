package queries

import "github.com/zizouhuweidi/retro-station/internal/pkg/utils"

// Ref: https://golangbot.com/inheritance/

type GetOrders struct {
	*utils.ListQuery
}

func NewGetOrders(query *utils.ListQuery) *GetOrders {
	return &GetOrders{ListQuery: query}
}
