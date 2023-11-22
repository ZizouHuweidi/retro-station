package dto

import (
	"time"
)

type GameDto struct {
	Id          string    `json:"id"`
	GameId      string    `json:"gameId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Genre       string    `json:"genre"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
