package data

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
)

type GameRepository interface {
	GetAllGames(
		ctx context.Context,
		listQuery *utils.ListQuery,
	) (*utils.ListResult[*models.Game], error)
	SearchGames(
		ctx context.Context,
		searchText string,
		listQuery *utils.ListQuery,
	) (*utils.ListResult[*models.Game], error)
	GetGameById(ctx context.Context, uuid string) (*models.Game, error)
	GetGameByGameId(ctx context.Context, uuid string) (*models.Game, error)
	CreateGame(ctx context.Context, game *models.Game) (*models.Game, error)
	UpdateGame(ctx context.Context, game *models.Game) (*models.Game, error)
	DeleteGameByID(ctx context.Context, uuid string) error
}
