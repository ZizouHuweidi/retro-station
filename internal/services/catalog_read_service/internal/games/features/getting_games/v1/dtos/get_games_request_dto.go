package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
)

type GetGamesRequestDto struct {
	*utils.ListQuery
}
