package queries

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	uuid "github.com/satori/go.uuid"
)

type GetGameById struct {
	Id uuid.UUID
}

func NewGetGameById(id uuid.UUID) (*GetGameById, error) {
	game := &GetGameById{Id: id}
	if err := game.Validate(); err != nil {
		return nil, err
	}

	return game, nil
}

func (p *GetGameById) Validate() error {
	return validation.ValidateStruct(p, validation.Field(&p.Id, validation.Required, is.UUIDv4))
}
