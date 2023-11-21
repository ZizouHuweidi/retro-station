package externalEvents

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"github.com/go-playground/validator"
	"github.com/mehdihadeli/go-mediatr"
	uuid "github.com/satori/go.uuid"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/consumer"
	messageTracing "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/updating_games/v1/commands"
)

type gameUpdatedConsumer struct {
	logger    logger.Logger
	validator *validator.Validate
	tracer    tracing.AppTracer
}

func NewGameUpdatedConsumer(
	logger logger.Logger,
	validator *validator.Validate,
	tracer tracing.AppTracer,
) consumer.ConsumerHandler {
	return &gameUpdatedConsumer{logger: logger, validator: validator, tracer: tracer}
}

func (c *gameUpdatedConsumer) Handle(
	ctx context.Context,
	consumeContext types.MessageConsumeContext,
) error {
	message, ok := consumeContext.Message().(*GameUpdatedV1)
	if !ok {
		return errors.New("error in casting message to GameUpdatedV1")
	}

	ctx, span := c.tracer.Start(ctx, "gameUpdatedConsumer.Handle")
	span.SetAttributes(attribute.Object("Message", consumeContext.Message()))
	defer span.End()

	gameUUID, err := uuid.FromString(message.GameId)
	if err != nil {
		c.logger.WarnMsg("uuid.FromString", err)
		badRequestErr := customErrors.NewBadRequestErrorWrap(
			err,
			"[updateGameConsumer_Consume.uuid.FromString] error in the converting uuid",
		)
		c.logger.Errorf(
			fmt.Sprintf(
				"[updateGameConsumer_Consume.uuid.FromString] err: %v",
				messageTracing.TraceMessagingErrFromSpan(span, badRequestErr),
			),
		)
		return err
	}

	command, err := commands.NewUpdateGame(
		gameUUID,
		message.Name,
		message.Description,
		message.Price,
		message.Genre,
	)
	if err != nil {
		validationErr := customErrors.NewValidationErrorWrap(
			err,
			"[updateGameConsumer_Consume.NewValidationErrorWrap] command validation failed",
		)
		c.logger.Errorf(
			fmt.Sprintf(
				"[updateGameConsumer_Consume.StructCtx] err: {%v}",
				messageTracing.TraceMessagingErrFromSpan(span, validationErr),
			),
		)
		return err
	}

	_, err = mediatr.Send[*commands.UpdateGame, *mediatr.Unit](ctx, command)
	if err != nil {
		err = errors.WithMessage(
			err,
			"[updateGameConsumer_Consume.Send] error in sending UpdateGame",
		)
		c.logger.Errorw(
			fmt.Sprintf(
				"[updateGameConsumer_Consume.Send] id: {%s}, err: {%v}",
				command.GameId,
				messageTracing.TraceMessagingErrFromSpan(span, err),
			),
			logger.Fields{"Id": command.GameId},
		)
		return err
	}

	return nil
}
