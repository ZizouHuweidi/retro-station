package commands

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	uuid "github.com/satori/go.uuid"
)

type CreateGame struct {
	// we generate id ourselves because auto generate mongo string id column with type _id is not an uuid
	Id          string
	GameId      string
	Name        string
	Description string
	Price       float64
	Genre       string
	CreatedAt   time.Time
}

func NewCreateGame(
	gameId string,
	name string,
	description string,
	price float64,
	genre string,
	createdAt time.Time,
) (*CreateGame, error) {
	command := &CreateGame{
		Id:          uuid.NewV4().String(),
		GameId:      gameId,
		Name:        name,
		Description: description,
		Price:       price,
		Genre:       genre,
		CreatedAt:   createdAt,
	}
	if err := command.Validate(); err != nil {
		return nil, err
	}

	return command, nil
}

func (g *CreateGame) Validate() error {
	return validation.ValidateStruct(g, validation.Field(&g.Id, validation.Required),
		validation.Field(&g.GameId, validation.Required),
		validation.Field(&g.Name, validation.Required, validation.Length(3, 250)),
		validation.Field(&g.Description, validation.Required, validation.Length(3, 500)),
		validation.Field(&g.Price, validation.Required),
		validation.Field(&g.Genre, validation.Required),
		validation.Field(&g.CreatedAt, validation.Required))
}
