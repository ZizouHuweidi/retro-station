package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"emperror.dev/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
)

const (
	redisGamePrefixKey = "game_read_service"
)

type redisGameRepository struct {
	log         logger.Logger
	redisClient redis.UniversalClient
	tracer      tracing.AppTracer
}

func NewRedisGameRepository(
	log logger.Logger,
	redisClient redis.UniversalClient,
	tracer tracing.AppTracer,
) data.GameCacheRepository {
	return &redisGameRepository{log: log, redisClient: redisClient, tracer: tracer}
}

func (r *redisGameRepository) PutGame(
	ctx context.Context,
	key string,
	game *models.Game,
) error {
	ctx, span := r.tracer.Start(ctx, "redisRepository.PutGame")
	span.SetAttributes(attribute2.String("PrefixKey", r.getRedisGamePrefixKey()))
	span.SetAttributes(attribute2.String("Key", key))
	defer span.End()

	gameBytes, err := json.Marshal(game)
	if err != nil {
		return tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				"[redisGameRepository_PutGame.Marshal] error marshalling game",
			),
		)
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisGamePrefixKey(), key, gameBytes).Err(); err != nil {
		return tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"[redisGameRepository_PutGame.HSetNX] error in updating game with key %s",
					key,
				),
			),
		)
	}

	span.SetAttributes(attribute.Object("Game", game))

	r.log.Infow(
		fmt.Sprintf(
			"[redisGameRepository.PutGame] game with key '%s', prefix '%s'  updated successfully",
			key,
			r.getRedisGamePrefixKey(),
		),
		logger.Fields{
			"Game":      game,
			"Id":        game.GameId,
			"Key":       key,
			"PrefixKey": r.getRedisGamePrefixKey(),
		},
	)

	return nil
}

func (r *redisGameRepository) GetGameById(
	ctx context.Context,
	key string,
) (*models.Game, error) {
	ctx, span := r.tracer.Start(ctx, "redisRepository.GetGameById")
	span.SetAttributes(attribute2.String("PrefixKey", r.getRedisGamePrefixKey()))
	span.SetAttributes(attribute2.String("Key", key))
	defer span.End()

	gameBytes, err := r.redisClient.HGet(ctx, r.getRedisGamePrefixKey(), key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"[redisGameRepository_GetGame.HGet] error in getting game with Key %s from database",
					key,
				),
			),
		)
	}

	var game models.Game
	if err := json.Unmarshal(gameBytes, &game); err != nil {
		return nil, tracing.TraceErrFromSpan(span, err)
	}

	span.SetAttributes(attribute.Object("Game", game))

	r.log.Infow(
		fmt.Sprintf(
			"[redisGameRepository.GetGameById] game with with key '%s', prefix '%s' laoded",
			key,
			r.getRedisGamePrefixKey(),
		),
		logger.Fields{
			"Game":      game,
			"Id":        game.GameId,
			"Key":       key,
			"PrefixKey": r.getRedisGamePrefixKey(),
		},
	)

	return &game, nil
}

func (r *redisGameRepository) DeleteGame(ctx context.Context, key string) error {
	ctx, span := r.tracer.Start(ctx, "redisRepository.DeleteGame")
	span.SetAttributes(attribute2.String("PrefixKey", r.getRedisGamePrefixKey()))
	span.SetAttributes(attribute2.String("Key", key))
	defer span.End()

	if err := r.redisClient.HDel(ctx, r.getRedisGamePrefixKey(), key).Err(); err != nil {
		return tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"[redisGameRepository_DeleteGame.HDel] error in deleting game with key %s",
					key,
				),
			),
		)
	}

	r.log.Infow(
		fmt.Sprintf(
			"[redisGameRepository.DeleteGame] game with key %s, prefix: %s deleted successfully",
			key,
			r.getRedisGamePrefixKey(),
		),
		logger.Fields{"Key": key, "PrefixKey": r.getRedisGamePrefixKey()},
	)

	return nil
}

func (r *redisGameRepository) DeleteAllGames(ctx context.Context) error {
	ctx, span := r.tracer.Start(ctx, "redisRepository.DeleteAllGames")
	span.SetAttributes(attribute2.String("PrefixKey", r.getRedisGamePrefixKey()))
	defer span.End()

	if err := r.redisClient.Del(ctx, r.getRedisGamePrefixKey()).Err(); err != nil {
		return tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				"[redisGameRepository_DeleteAllGames.Del] error in deleting all games",
			),
		)
	}

	r.log.Infow(
		"[redisGameRepository.DeleteAllGames] all games deleted",
		logger.Fields{"PrefixKey": r.getRedisGamePrefixKey()},
	)

	return nil
}

func (r *redisGameRepository) getRedisGamePrefixKey() string {
	return redisGamePrefixKey
}
