package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
)

type SearchGamesResponseDto struct {
	Games *utils.ListResult[*dtoV1.GameDto]
}
