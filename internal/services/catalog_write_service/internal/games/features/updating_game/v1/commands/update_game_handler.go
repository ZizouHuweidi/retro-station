package commands

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mehdihadeli/go-mediatr"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/producer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	dto "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/events/integration_events"
)

type UpdateGameHandler struct {
	log              logger.Logger
	uow              data.CatalogUnitOfWork
	rabbitmqProducer producer.Producer
	tracer           tracing.AppTracer
}

func NewUpdateGameHandler(
	log logger.Logger,
	uow data.CatalogUnitOfWork,
	rabbitmqProducer producer.Producer,
	tracer tracing.AppTracer,
) *UpdateGameHandler {
	return &UpdateGameHandler{
		log:              log,
		uow:              uow,
		rabbitmqProducer: rabbitmqProducer,
		tracer:           tracer,
	}
}

func (c *UpdateGameHandler) Handle(
	ctx context.Context,
	command *UpdateGame,
) (*mediatr.Unit, error) {
	ctx, span := c.tracer.Start(ctx, "UpdateGameHandler.Handle")
	span.SetAttributes(attribute2.String("GameId", command.GameID.String()))
	span.SetAttributes(attribute.Object("Command", command))
	defer span.End()

	err := c.uow.Do(ctx, func(catalogContext data.CatalogContext) error {
		game, err := catalogContext.Games().GetGameById(ctx, command.GameID)
		if err != nil {
			return tracing.TraceErrFromSpan(
				span,
				customErrors.NewApplicationErrorWrapWithCode(
					err,
					http.StatusNotFound,
					fmt.Sprintf(
						"[UpdateGameHandler_Handle.GetGameById] game with id %s not found",
						command.GameID,
					),
				),
			)
		}

		game.Name = command.Name
		game.Price = command.Price
		game.Description = command.Description
		game.Genre = command.Genre
		game.UpdatedAt = command.UpdatedAt

		updatedGame, err := catalogContext.Games().UpdateGame(ctx, game)
		if err != nil {
			return tracing.TraceErrFromSpan(
				span,
				customErrors.NewApplicationErrorWrap(
					err,
					"[UpdateGameHandler_Handle.UpdateGame] error in updating game in the repository",
				),
			)
		}

		gameDto, err := mapper.Map[*dto.GameDto](updatedGame)
		if err != nil {
			return tracing.TraceErrFromSpan(
				span,
				customErrors.NewApplicationErrorWrap(
					err,
					"[UpdateGameHandler_Handle.Map] error in the mapping GameDto",
				),
			)
		}

		gameUpdated := integration_events.NewGameUpdatedV1(gameDto)

		err = c.rabbitmqProducer.PublishMessage(ctx, gameUpdated, nil)
		if err != nil {
			return tracing.TraceErrFromSpan(
				span,
				customErrors.NewApplicationErrorWrap(
					err,
					"[UpdateGameHandler_Handle.PublishMessage] error in publishing 'GameUpdated' message",
				),
			)
		}

		c.log.Infow(
			fmt.Sprintf(
				"[UpdateGameHandler.Handle] game with id '%s' updated",
				command.GameID,
			),
			logger.Fields{"GameId": command.GameID},
		)

		c.log.Infow(
			fmt.Sprintf(
				"[DeleteGameHandler.Handle] GameUpdated message with messageId `%s` published to the rabbitmq broker",
				gameUpdated.MessageId,
			),
			logger.Fields{"MessageId": gameUpdated.MessageId},
		)

		return nil
	})

	return &mediatr.Unit{}, err
}
