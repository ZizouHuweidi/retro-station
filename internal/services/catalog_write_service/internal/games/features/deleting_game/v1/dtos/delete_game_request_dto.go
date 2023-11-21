package dtos

import (
	uuid "github.com/satori/go.uuid"
)

type DeleteGameRequestDto struct {
	GameID uuid.UUID `param:"id" json:"-"`
}
