package dtos

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
)

// https://echo.labstack.com/guide/binding/
// https://echo.labstack.com/guide/request/
// https://github.com/go-playground/validator

// GetGamesRequestDto validation will handle in command level
type GetGamesRequestDto struct {
	*utils.ListQuery
}
