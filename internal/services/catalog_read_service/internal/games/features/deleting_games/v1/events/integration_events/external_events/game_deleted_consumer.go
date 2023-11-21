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

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/deleting_games/v1/commands"
)

type gameDeletedConsumer struct {
	logger    logger.Logger
	validator *validator.Validate
	tracer    tracing.AppTracer
}

func NewGameDeletedConsumer(
	logger logger.Logger,
	validator *validator.Validate,
	tracer tracing.AppTracer,
) consumer.ConsumerHandler {
	return &gameDeletedConsumer{logger: logger, validator: validator, tracer: tracer}
}

func (c *gameDeletedConsumer) Handle(
	ctx context.Context,
	consumeContext types.MessageConsumeContext,
) error {
	message, ok := consumeContext.Message().(*GameDeletedV1)
	if !ok {
		return errors.New("error in casting message to GameDeletedV1")
	}

	ctx, span := c.tracer.Start(ctx, "gameDeletedConsumer.Handle")
	span.SetAttributes(attribute.Object("Message", consumeContext.Message()))
	defer span.End()

	gameUUID, err := uuid.FromString(message.GameId)
	if err != nil {
		badRequestErr := customErrors.NewBadRequestErrorWrap(
			err,
			"[gameDeletedConsumer_Handle.uuid.FromString] error in the converting uuid",
		)
		c.logger.Errorf(
			fmt.Sprintf(
				"[gameDeletedConsumer_Handle.uuid.FromString] err: %v",
				messageTracing.TraceMessagingErrFromSpan(span, badRequestErr),
			),
		)

		return err
	}

	command, err := commands.NewDeleteGame(gameUUID)
	if err != nil {
		validationErr := customErrors.NewValidationErrorWrap(
			err,
			"[gameDeletedConsumer_Handle.StructCtx] command validation failed",
		)
		c.logger.Errorf(
			fmt.Sprintf(
				"[gameDeletedConsumer_Consume.StructCtx] err: {%v}",
				messageTracing.TraceMessagingErrFromSpan(span, validationErr),
			),
		)

		return err
	}

	_, err = mediatr.Send[*commands.DeleteGame, *mediatr.Unit](ctx, command)

	return err
}
