package createGameCommand

import (
	"context"
	"fmt"
	"net/http"

	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/producer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/dtos"
	integrationEvents "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/events/integration_events"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
)

type CreateGameHandler struct {
	log              logger.Logger
	uow              data.CatalogUnitOfWork
	rabbitmqProducer producer.Producer
	tracer           tracing.AppTracer
}

func NewCreateGameHandler(
	log logger.Logger,
	uow data.CatalogUnitOfWork,
	rabbitmqProducer producer.Producer,
	tracer tracing.AppTracer,
) *CreateGameHandler {
	return &CreateGameHandler{
		log:              log,
		uow:              uow,
		rabbitmqProducer: rabbitmqProducer,
		tracer:           tracer,
	}
}

func (c *CreateGameHandler) Handle(
	ctx context.Context,
	command *CreateGame,
) (*dtos.CreateGameResponseDto, error) {
	ctx, span := c.tracer.Start(ctx, "CreateGameHandler.Handle")
	span.SetAttributes(attribute2.String("GameId", command.GameID.String()))
	span.SetAttributes(attribute.Object("Command", command))
	defer span.End()

	game := &models.Game{
		GameId:      command.GameID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		Genre:       command.Genre,
		CreatedAt:   command.CreatedAt,
	}

	var createGameResult *dtos.CreateGameResponseDto

	err := c.uow.Do(ctx, func(catalogContext data.CatalogContext) error {
		createdGame, err := catalogContext.Games().CreateGame(ctx, game)
		if err != nil {
			return tracing.TraceErrFromSpan(
				span,
				customErrors.NewApplicationErrorWrapWithCode(
					err,
					http.StatusConflict,
					"[CreateGameHandler.CreateGame] game already exists",
				),
			)
		}
		gameDto, err := mapper.Map[*dtoV1.GameDto](createdGame)
		if err != nil {
			return tracing.TraceErrFromSpan(
				span,
				customErrors.NewApplicationErrorWrap(
					err,
					"[CreateGameHandler.Map] error in the mapping GameDto",
				),
			)
		}

		gameCreated := integrationEvents.NewGameCreatedV1(gameDto)

		err = c.rabbitmqProducer.PublishMessage(ctx, gameCreated, nil)
		if err != nil {
			return tracing.TraceErrFromSpan(
				span,
				customErrors.NewApplicationErrorWrap(
					err,
					"[CreateGameHandler.PublishMessage] error in publishing GameCreated integration_events event",
				),
			)
		}

		c.log.Infow(
			fmt.Sprintf(
				"[CreateGameHandler.Handle] GameCreated message with messageId `%s` published to the rabbitmq broker",
				gameCreated.MessageId,
			),
			logger.Fields{"MessageId": gameCreated.MessageId},
		)

		createGameResult = &dtos.CreateGameResponseDto{GameID: game.GameId}

		span.SetAttributes(attribute.Object("CreateGameResultDto", createGameResult))

		c.log.Infow(
			fmt.Sprintf(
				"[CreateGameHandler.Handle] game with id '%s' created",
				command.GameID,
			),
			logger.Fields{"GameId": command.GameID, "MessageId": gameCreated.MessageId},
		)

		return nil
	})

	return createGameResult, err
}
