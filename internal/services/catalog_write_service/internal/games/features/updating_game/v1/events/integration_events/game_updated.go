package integration_events

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"

	dto "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
)

type GameUpdatedV1 struct {
	*types.Message
	*dto.GameDto
}

func NewGameUpdatedV1(gameDto *dto.GameDto) *GameUpdatedV1 {
	return &GameUpdatedV1{
		Message: types.NewMessage(uuid.NewV4().String()),
		GameDto: gameDto,
	}
}
