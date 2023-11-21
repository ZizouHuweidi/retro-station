package queries

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
)

type SearchGames struct {
	SearchText string
	*utils.ListQuery
}

func NewSearchGames(searchText string, query *utils.ListQuery) (*SearchGames, error) {
	command := &SearchGames{
		SearchText: searchText,
		ListQuery:  query,
	}

	err := command.Validate()
	if err != nil {
		return nil, err
	}

	return command, nil
}

func (p *SearchGames) Validate() error {
	return validation.ValidateStruct(p, validation.Field(&p.SearchText, validation.Required))
}
