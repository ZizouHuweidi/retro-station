package data

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
)

type GameCacheRepository interface {
	PutGame(ctx context.Context, key string, game *models.Game) error
	GetGameById(ctx context.Context, key string) (*models.Game, error)
	DeleteGame(ctx context.Context, key string) error
	DeleteAllGames(ctx context.Context) error
}
