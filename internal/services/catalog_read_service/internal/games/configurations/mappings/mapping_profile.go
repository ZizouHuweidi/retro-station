package mappings

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/dto"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
)

func ConfigureGamesMappings() error {
	err := mapper.CreateMap[*models.Game, *dto.GameDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*models.Game, *models.Game]()
	if err != nil {
		return err
	}

	return nil
}
