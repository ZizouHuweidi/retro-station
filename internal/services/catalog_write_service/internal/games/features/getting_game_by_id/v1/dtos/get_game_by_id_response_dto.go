package dtos

import dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"

// https://echo.labstack.com/guide/response/
type GetGameByIdResponseDto struct {
	Game *dtoV1.GameDto `json:"game"`
}
