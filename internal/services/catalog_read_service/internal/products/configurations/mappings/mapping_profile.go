package mappings

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/products/dto"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/products/models"
)

func ConfigureProductsMappings() error {
	err := mapper.CreateMap[*models.Product, *dto.ProductDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*models.Product, *models.Product]()
	if err != nil {
		return err
	}

	return nil
}
