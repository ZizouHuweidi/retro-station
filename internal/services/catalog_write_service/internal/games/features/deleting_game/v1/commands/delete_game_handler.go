package commands

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mehdihadeli/go-mediatr"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/producer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	integrationEvents "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/deleting_game/v1/events/integration_events"
)

type DeleteGameHandler struct {
	log              logger.Logger
	uow              data.CatalogUnitOfWork
	rabbitmqProducer producer.Producer
	tracer           tracing.AppTracer
}

func NewDeleteGameHandler(
	log logger.Logger,
	uow data.CatalogUnitOfWork,
	rabbitmqProducer producer.Producer,
	tracer tracing.AppTracer,
) *DeleteGameHandler {
	return &DeleteGameHandler{
		log:              log,
		uow:              uow,
		rabbitmqProducer: rabbitmqProducer,
		tracer:           tracer,
	}
}

func (c *DeleteGameHandler) Handle(
	ctx context.Context,
	command *DeleteGame,
) (*mediatr.Unit, error) {
	ctx, span := c.tracer.Start(ctx, "deleteGameHandler.Handle")
	span.SetAttributes(attribute2.String("GameId", command.GameID.String()))
	span.SetAttributes(attribute.Object("Command", command))
	defer span.End()

	err := c.uow.Do(ctx, func(catalogContext data.CatalogContext) error {
		if err := catalogContext.Games().DeleteGameByID(ctx, command.GameID); err != nil {
			return tracing.TraceErrFromSpan(
				span,
				customErrors.NewApplicationErrorWrapWithCode(
					err,
					http.StatusNotFound,
					"[DeleteGameHandler_Handle.DeleteGameByID] game not found",
				),
			)
		}

		gameDeleted := integrationEvents.NewGameDeletedV1(command.GameID.String())
		err := c.rabbitmqProducer.PublishMessage(ctx, gameDeleted, nil)
		if err != nil {
			return tracing.TraceErrFromSpan(
				span,
				customErrors.NewApplicationErrorWrap(
					err,
					"[DeleteGameHandler_Handle.PublishMessage] error in publishing 'GameDeleted' message",
				),
			)
		}

		c.log.Infow(
			fmt.Sprintf(
				"[DeleteGameHandler.Handle] GameDeleted message with messageId '%s' published to the rabbitmq broker",
				gameDeleted.MessageId,
			),
			logger.Fields{"MessageId": gameDeleted.MessageId},
		)

		c.log.Infow(
			fmt.Sprintf(
				"[DeleteGameHandler.Handle] game with id '%s' deleted",
				command.GameID,
			),
			logger.Fields{"GameId": command.GameID},
		)

		return nil
	})

	return &mediatr.Unit{}, err
}
