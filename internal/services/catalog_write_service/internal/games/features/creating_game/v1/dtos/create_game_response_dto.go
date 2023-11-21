package dtos

import (
	uuid "github.com/satori/go.uuid"
)

// https://echo.labstack.com/guide/response/
type CreateGameResponseDto struct {
	GameID uuid.UUID `json:"gameId"`
}
