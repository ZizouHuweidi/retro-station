package queries

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
)

type SearchGames struct {
	SearchText string
	*utils.ListQuery
}

func (s *SearchGames) Validate() error {
	return validation.ValidateStruct(s, validation.Field(&s.SearchText, validation.Required))
}
