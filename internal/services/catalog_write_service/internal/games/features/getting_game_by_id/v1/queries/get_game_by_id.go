package getGameByIdQuery

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	uuid "github.com/satori/go.uuid"
)

// https://echo.labstack.com/guide/request/
// https://github.com/go-playground/validator

type GetGameById struct {
	GameID uuid.UUID
}

func NewGetGameById(gameId uuid.UUID) (*GetGameById, error) {
	query := &GetGameById{GameID: gameId}
	err := query.Validate()
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (p *GetGameById) Validate() error {
	return validation.ValidateStruct(p, validation.Field(&p.GameID, validation.Required, is.UUIDv4))
}
