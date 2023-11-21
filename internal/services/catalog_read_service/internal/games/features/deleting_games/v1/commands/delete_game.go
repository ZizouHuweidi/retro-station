package commands

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	uuid "github.com/satori/go.uuid"
)

type DeleteGame struct {
	GameId uuid.UUID
}

func NewDeleteGame(gameId uuid.UUID) (*DeleteGame, error) {
	delGame := &DeleteGame{GameId: gameId}
	if err := delGame.Validate(); err != nil {
		return nil, err
	}

	return delGame, nil
}

func (p *DeleteGame) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.GameId, validation.Required, is.UUIDv4))
}
