package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
)

type SearchProductsRequestDto struct {
	SearchText       string `query:"search" json:"search"`
	*utils.ListQuery `                      json:"listQuery"`
}
