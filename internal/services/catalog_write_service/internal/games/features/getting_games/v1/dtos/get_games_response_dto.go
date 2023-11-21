package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
)

// https://echo.labstack.com/guide/response/
type GetGamesResponseDto struct {
	Games *utils.ListResult[*dtoV1.GameDto]
}
