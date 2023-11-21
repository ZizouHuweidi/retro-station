package commands

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	uuid "github.com/satori/go.uuid"
)

type UpdateGame struct {
	GameId      uuid.UUID
	Name        string
	Description string
	Price       float64
	Genre       string
	UpdatedAt   time.Time
}

func NewUpdateGame(gameId uuid.UUID, name string, description string, price float64, genre string) (*UpdateGame, error) {
	game := &UpdateGame{
		GameId:      gameId,
		Name:        name,
		Description: description,
		Price:       price,
		Genre:       genre,
		UpdatedAt:   time.Now(),
	}
	if err := game.Validate(); err != nil {
		return nil, err
	}
	return game, nil
}

func (p *UpdateGame) Validate() error {
	return validation.ValidateStruct(p, validation.Field(&p.GameId, validation.Required, is.UUIDv4),
		validation.Field(&p.Name, validation.Required, validation.Length(0, 255)),
		validation.Field(&p.Description, validation.Required, validation.Length(0, 5000)),
		validation.Field(&p.Price, validation.Required, validation.Min(0.0)),
		validation.Field(&p.UpdatedAt, validation.Required),
	)
}
