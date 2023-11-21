package integration_events

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"

	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
)

type GameCreatedV1 struct {
	*types.Message
	*dtoV1.GameDto
}

func NewGameCreatedV1(gameDto *dtoV1.GameDto) *GameCreatedV1 {
	return &GameCreatedV1{
		GameDto: gameDto,
		Message: types.NewMessage(uuid.NewV4().String()),
	}
}
