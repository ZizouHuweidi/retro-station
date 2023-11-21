package integrationEvents

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
)

type GameDeletedV1 struct {
	*types.Message
	GameId string `json:"gameId,omitempty"`
}

func NewGameDeletedV1(gameId string) *GameDeletedV1 {
	return &GameDeletedV1{GameId: gameId, Message: types.NewMessage(uuid.NewV4().String())}
}
