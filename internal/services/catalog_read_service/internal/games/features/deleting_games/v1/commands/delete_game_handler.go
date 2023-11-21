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

type DeleteGameCommand struct {
	log             logger.Logger
	mongoRepository data.GameRepository
	redisRepository data.GameCacheRepository
	tracer          tracing.AppTracer
}

func NewDeleteGameHandler(
	log logger.Logger,
	repository data.GameRepository,
	redisRepository data.GameCacheRepository,
	tracer tracing.AppTracer,
) *DeleteGameCommand {
	return &DeleteGameCommand{
		log:             log,
		mongoRepository: repository,
		redisRepository: redisRepository,
		tracer:          tracer,
	}
}

func (c *DeleteGameCommand) Handle(
	ctx context.Context,
	command *DeleteGame,
) (*mediatr.Unit, error) {
	ctx, span := c.tracer.Start(ctx, "DeleteGameCommand.Handle")
	span.SetAttributes(attribute2.String("GameId", command.GameId.String()))
	span.SetAttributes(attribute.Object("Command", command))
	game, err := c.mongoRepository.GetGameByGameId(ctx, command.GameId.String())

	defer span.End()

	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				fmt.Sprintf(
					"[DeleteGameHandler_Handle.GetGameById] error in fetching game with gameId %s in the mongo repository",
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
					"[DeleteGameHandler_Handle.GetGameById] game with gameId %s not found",
					command.GameId,
				),
			),
		)
	}

	if err := c.mongoRepository.DeleteGameByID(ctx, game.Id); err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[DeleteGameHandler_Handle.DeleteGameByID] error in deleting game in the mongo repository",
			),
		)
	}

	c.log.Infof("(game deleted) id: {%s}", game.Id)

	err = c.redisRepository.DeleteGame(ctx, game.Id)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[DeleteGameHandler_Handle.DeleteGame] error in deleting game in the redis repository",
			),
		)
	}

	c.log.Infow(
		fmt.Sprintf("[DeleteGameCommand.Handle] game with id: {%s} deleted", game.Id),
		logger.Fields{"GameId": command.GameId, "Id": game.Id},
	)

	return &mediatr.Unit{}, nil
}
