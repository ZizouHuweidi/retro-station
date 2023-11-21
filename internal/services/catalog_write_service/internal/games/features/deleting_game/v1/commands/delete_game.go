package commands

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	uuid "github.com/satori/go.uuid"
)

type DeleteGame struct {
	GameID uuid.UUID
}

func NewDeleteGame(gameID uuid.UUID) (*DeleteGame, error) {
	command := &DeleteGame{GameID: gameID}
	err := command.Validate()
	if err != nil {
		return nil, err
	}

	return command, nil
}

func (p *DeleteGame) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.GameID, validation.Required),
		validation.Field(&p.GameID, is.UUIDv4),
	)
}
