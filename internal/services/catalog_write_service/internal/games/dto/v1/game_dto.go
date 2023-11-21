package dtoV1

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type GameDto struct {
	GameId      uuid.UUID `json:"gameId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Genre       string    `json:"genre"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
