package commands

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	uuid "github.com/satori/go.uuid"
)

type UpdateGame struct {
	GameID      uuid.UUID
	Name        string
	Description string
	Price       float64
	Genre       string
	UpdatedAt   time.Time
}

func NewUpdateGame(gameID uuid.UUID, name string, description string, price float64, genre string) (*UpdateGame, error) {
	command := &UpdateGame{
		GameID:      gameID,
		Name:        name,
		Description: description,
		Price:       price,
		Genre:       genre,
		UpdatedAt:   time.Now(),
	}
	err := command.Validate()
	if err != nil {
		return nil, err
	}

	return command, nil
}

func (p *UpdateGame) Validate() error {
	return validation.ValidateStruct(p, validation.Field(&p.GameID, validation.Required),
		validation.Field(&p.Name, validation.Required, validation.Length(0, 255)),
		validation.Field(&p.Description, validation.Required, validation.Length(0, 5000)),
		validation.Field(&p.Price, validation.Required, validation.Min(0.0)),
		validation.Field(&p.Genre, validation.Required, validation.Length(0, 255)),
		validation.Field(&p.UpdatedAt, validation.Required))
}
