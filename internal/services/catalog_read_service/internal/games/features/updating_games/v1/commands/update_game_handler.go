package commands

import (
	"context"
	"fmt"

	"github.com/mehdihadeli/go-mediatr"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/data"
)

type UpdateGameHandler struct {
	log             logger.Logger
	mongoRepository data.GameRepository
	redisRepository data.GameCacheRepository
	tracer          tracing.AppTracer
}

func NewUpdateGameHandler(
	log logger.Logger,
	mongoRepository data.GameRepository,
	redisRepository data.GameCacheRepository,
	tracer tracing.AppTracer,
) *UpdateGameHandler {
	return &UpdateGameHandler{
		log:             log,
		mongoRepository: mongoRepository,
		redisRepository: redisRepository,
		tracer:          tracer,
	}
}

func (c *UpdateGameHandler) Handle(
	ctx context.Context,
	command *UpdateGame,
) (*mediatr.Unit, error) {
	ctx, span := c.tracer.Start(ctx, "UpdateGameHandler.Handle")
	span.SetAttributes(attribute2.String("GameId", command.GameId.String()))
	span.SetAttributes(attribute.Object("Command", command))
	defer span.End()

	game, err := c.mongoRepository.GetGameByGameId(ctx, command.GameId.String())
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				fmt.Sprintf(
					"[UpdateGameHandler_Handle.GetGameById] error in fetching game with gameId %s in the mongo repository",
					command.GameId,
				),
			),
		)
	}

	if game == nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewNotFoundErrorWrap(
				err,
				fmt.Sprintf(
					"[UpdateGameHandler_Handle.GetGameById] game with gameId %s not found",
					command.GameId,
				),
			),
		)
	}

	game.Price = command.Price
	game.Name = command.Name
	game.Description = command.Description
	game.UpdatedAt = command.UpdatedAt

	_, err = c.mongoRepository.UpdateGame(ctx, game)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[UpdateGameHandler_Handle.UpdateGame] error in updating game in the mongo repository",
			),
		)
	}

	err = c.redisRepository.PutGame(ctx, game.Id, game)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[UpdateGameHandler_Handle.PutGame] error in updating game in the redis repository",
			),
		)
	}

	c.log.Infow(
		fmt.Sprintf("[UpdateGameHandler.Handle] game with id: {%s} updated", game.Id),
		logger.Fields{"GameId": command.GameId, "Id": game.Id},
	)

	return &mediatr.Unit{}, nil
}
