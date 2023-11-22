package externalEvents

import (
	"time"

	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
)

type GameCreatedV1 struct {
	*types.Message
	GameId      string    `json:"gameId,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}
