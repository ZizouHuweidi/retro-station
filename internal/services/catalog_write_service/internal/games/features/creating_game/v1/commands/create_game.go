package createGameCommand

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	uuid "github.com/satori/go.uuid"
)

// https://echo.labstack.com/guide/request/
// https://github.com/go-playground/validator

type CreateGame struct {
	GameID      uuid.UUID
	Name        string
	Description string
	Price       float64
	Genre       string
	CreatedAt   time.Time
}

func NewCreateGame(name string, description string, price float64, genre string) (*CreateGame, error) {
	command := &CreateGame{
		GameID:      uuid.NewV4(),
		Name:        name,
		Description: description,
		Price:       price,
		Genre:       genre,
		CreatedAt:   time.Now(),
	}
	err := command.Validate()
	if err != nil {
		return nil, err
	}
	return command, nil
}

func (c *CreateGame) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.GameID, validation.Required),
		validation.Field(&c.Name, validation.Required, validation.Length(0, 255)),
		validation.Field(&c.Description, validation.Required, validation.Length(0, 5000)),
		validation.Field(&c.Price, validation.Required, validation.Min(0.0).Exclusive()),
		validation.Field(&c.Genre, validation.Required, validation.Length(0, 255)),
		validation.Field(&c.CreatedAt, validation.Required),
	)
}
