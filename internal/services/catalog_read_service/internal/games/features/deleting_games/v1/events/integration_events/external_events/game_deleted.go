package externalEvents

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
)

type GameDeletedV1 struct {
	*types.Message
	GameId string `json:"gameId,omitempty"`
}
