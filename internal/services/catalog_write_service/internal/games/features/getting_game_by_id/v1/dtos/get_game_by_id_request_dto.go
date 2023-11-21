package dtos

import (
	uuid "github.com/satori/go.uuid"
)

// https://echo.labstack.com/guide/binding/
// https://echo.labstack.com/guide/request/
// https://github.com/go-playground/validator

// GetGameByIdRequestDto validation will handle in query level
type GetGameByIdRequestDto struct {
	GameId uuid.UUID `param:"id" json:"-"`
}
