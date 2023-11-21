package commands

import (
	"context"
	"fmt"

	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/creating_game/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
)

type CreateGameHandler struct {
	log             logger.Logger
	mongoRepository data.GameRepository
	redisRepository data.GameCacheRepository
	tracer          tracing.AppTracer
}

func NewCreateGameHandler(
	log logger.Logger,
	mongoRepository data.GameRepository,
	redisRepository data.GameCacheRepository,
	tracer tracing.AppTracer,
) *CreateGameHandler {
	return &CreateGameHandler{
		log:             log,
		mongoRepository: mongoRepository,
		redisRepository: redisRepository,
		tracer:          tracer,
	}
}

func (c *CreateGameHandler) Handle(
	ctx context.Context,
	command *CreateGame,
) (*dtos.CreateGameResponseDto, error) {
	ctx, span := c.tracer.Start(ctx, "CreateGameHandler.Handle")
	span.SetAttributes(attribute2.String("GameId", command.GameId))
	span.SetAttributes(attribute.Object("Command", command))

	defer span.End()

	game := &models.Game{
		Id:          command.Id, // we generate id ourselves because auto generate mongo string id column with type _id is not an uuid
		GameId:      command.GameId,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		Genre:       command.Genre,
		CreatedAt:   command.CreatedAt,
	}

	createdGame, err := c.mongoRepository.CreateGame(ctx, game)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[CreateGameHandler_Handle.CreateGame] error in creating game in the mongo repository",
			),
		)
	}

	err = c.redisRepository.PutGame(ctx, createdGame.Id, createdGame)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[CreateGameHandler_Handle.PutGame] error in creating game in the redis repository",
			),
		)
	}

	response := &dtos.CreateGameResponseDto{Id: createdGame.Id}
	span.SetAttributes(attribute.Object("CreateGameResponseDto", response))

	c.log.Infow(
		fmt.Sprintf("[CreateGameHandler.Handle] game with id: {%s} created", game.Id),
		logger.Fields{"GameId": command.GameId, "Id": game.Id},
	)

	return response, nil
}
