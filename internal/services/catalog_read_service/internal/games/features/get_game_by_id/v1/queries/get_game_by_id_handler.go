package queries

import (
	"context"
	"fmt"

	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/dto"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/get_game_by_id/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
)

type GetGameByIdHandler struct {
	log             logger.Logger
	mongoRepository data.GameRepository
	redisRepository data.GameCacheRepository
	tracer          tracing.AppTracer
}

func NewGetGameByIdHandler(
	log logger.Logger,
	mongoRepository data.GameRepository,
	redisRepository data.GameCacheRepository,
	tracer tracing.AppTracer,
) *GetGameByIdHandler {
	return &GetGameByIdHandler{
		log:             log,
		mongoRepository: mongoRepository,
		redisRepository: redisRepository,
		tracer:          tracer,
	}
}

func (q *GetGameByIdHandler) Handle(
	ctx context.Context,
	query *GetGameById,
) (*dtos.GetGameByIdResponseDto, error) {
	ctx, span := q.tracer.Start(ctx, "getGameByIdHandler.Handle")
	span.SetAttributes(attribute.Object("Query", query))
	span.SetAttributes(attribute2.String("Id", query.Id.String()))

	redisGame, err := q.redisRepository.GetGameById(ctx, query.Id.String())

	var game *models.Game

	defer span.End()

	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				fmt.Sprintf(
					"[GetGameByIdHandler_Handle.GetGameById] error in getting game with id %d in the redis repository",
					query.Id,
				),
			),
		)
	}

	if redisGame != nil {
		game = redisGame
	} else {
		var mongoGame *models.Game
		mongoGame, err = q.mongoRepository.GetGameById(ctx, query.Id.String())
		if err != nil {
			return nil, tracing.TraceErrFromSpan(span, customErrors.NewApplicationErrorWrap(err, fmt.Sprintf("[GetGameByIdHandler_Handle.GetGameById] error in getting game with id %d in the mongo repository", query.Id)))
		}
		if mongoGame == nil {
			mongoGame, err = q.mongoRepository.GetGameByGameId(ctx, query.Id.String())
		}
		if err != nil {
			return nil, err
		}

		game = mongoGame
		err = q.redisRepository.PutGame(ctx, game.Id, game)
		if err != nil {
			return new(dtos.GetGameByIdResponseDto), tracing.TraceErrFromSpan(span, err)
		}
	}

	gameDto, err := mapper.Map[*dto.GameDto](game)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[GetGameByIdHandler_Handle.Map] error in the mapping game",
			),
		)
	}

	q.log.Infow(
		fmt.Sprintf("[GetGameByIdHandler.Handle] game with id: {%s} fetched", query.Id),
		logger.Fields{"GameId": game.GameId, "Id": game.Id},
	)

	return &dtos.GetGameByIdResponseDto{Game: gameDto}, nil
}
