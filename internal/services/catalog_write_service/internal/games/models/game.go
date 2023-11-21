package models

import (
	"time"

	"github.com/goccy/go-json"
	uuid "github.com/satori/go.uuid"
)

// Game model
type Game struct {
	GameId      uuid.UUID `json:"gameId"   gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Genre       string    `json:"genre"`
	CreatedAt   time.Time `json:"createdAt"` // https://gorm.io/docs/models.html#gorm-Model
	UpdatedAt   time.Time `json:"updatedAt"` // https://gorm.io/docs/models.html#gorm-Model
}

func (p *Game) String() string {
	j, _ := json.Marshal(p)
	return string(j)
}
