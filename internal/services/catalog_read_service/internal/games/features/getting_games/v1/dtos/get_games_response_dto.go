package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/dto"
)

type GetGamesResponseDto struct {
	Games *utils.ListResult[*dto.GameDto]
}