package mappings

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	"google.golang.org/protobuf/types/known/timestamppb"

	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	gamesService "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc/genproto"
)

func ConfigureGamesMappings() error {
	err := mapper.CreateMap[*models.Game, *dtoV1.GameDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*dtoV1.GameDto, *models.Game]()
	if err != nil {
		return err
	}

	err = mapper.CreateCustomMap[*dtoV1.GameDto, *gamesService.Game](
		func(game *dtoV1.GameDto) *gamesService.Game {
			if game == nil {
				return nil
			}
			return &gamesService.Game{
				GameId:      game.GameId.String(),
				Name:        game.Name,
				Description: game.Description,
				Price:       game.Price,
				Genre:       game.Genre,
				CreatedAt:   timestamppb.New(game.CreatedAt),
				UpdatedAt:   timestamppb.New(game.UpdatedAt),
			}
		},
	)
	if err != nil {
		return err
	}

	err = mapper.CreateCustomMap(func(game *models.Game) *gamesService.Game {
		return &gamesService.Game{
			GameId:      game.GameId.String(),
			Name:        game.Name,
			Description: game.Description,
			Price:       game.Price,
			Genre:       game.Genre,
			CreatedAt:   timestamppb.New(game.CreatedAt),
			UpdatedAt:   timestamppb.New(game.UpdatedAt),
		}
	})

	return nil
}
