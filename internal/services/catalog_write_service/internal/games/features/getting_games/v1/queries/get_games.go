package queries

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
)

// Ref: https://golangbot.com/inheritance/

type GetGames struct {
	*utils.ListQuery
}

func NewGetGames(query *utils.ListQuery) (*GetGames, error) {
	return &GetGames{ListQuery: query}, nil
}
