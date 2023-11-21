package data

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
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
	GetGameById(ctx context.Context, uuid uuid.UUID) (*models.Game, error)
	CreateGame(ctx context.Context, game *models.Game) (*models.Game, error)
	UpdateGame(ctx context.Context, game *models.Game) (*models.Game, error)
	DeleteGameByID(ctx context.Context, uuid uuid.UUID) error
}
