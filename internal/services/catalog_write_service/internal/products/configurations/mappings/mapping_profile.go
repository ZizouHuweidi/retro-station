package mappings

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/dto/v1"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/models"
	productsService "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc/genproto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConfigureProductsMappings() error {
	err := mapper.CreateMap[*models.Product, *dtoV1.ProductDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*dtoV1.ProductDto, *models.Product]()
	if err != nil {
		return err
	}

	err = mapper.CreateCustomMap[*dtoV1.ProductDto, *productsService.Product](
		func(product *dtoV1.ProductDto) *productsService.Product {
			if product == nil {
				return nil
			}
			return &productsService.Product{
				ProductId:   product.ProductId.String(),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				CreatedAt:   timestamppb.New(product.CreatedAt),
				UpdatedAt:   timestamppb.New(product.UpdatedAt),
			}
		},
	)
	if err != nil {
		return err
	}

	err = mapper.CreateCustomMap(func(product *models.Product) *productsService.Product {
		return &productsService.Product{
			ProductId:   product.ProductId.String(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		}
	})

	return nil
}
