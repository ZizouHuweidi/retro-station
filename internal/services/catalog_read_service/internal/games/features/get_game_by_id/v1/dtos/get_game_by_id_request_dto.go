package dtos

import (
	uuid "github.com/satori/go.uuid"
)

type GetGameByIdRequestDto struct {
	Id uuid.UUID `param:"id" json:"-"`
}
