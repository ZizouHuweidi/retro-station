package dtos

import "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/dto"

type GetGameByIdResponseDto struct {
	Game *dto.GameDto `json:"game"`
}
