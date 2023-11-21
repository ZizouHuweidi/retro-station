package externalEvents

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"github.com/go-playground/validator"
	"github.com/mehdihadeli/go-mediatr"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/consumer"
	messageTracing "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/creating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/creating_game/v1/dtos"
)

type gameCreatedConsumer struct {
	logger    logger.Logger
	validator *validator.Validate
	tracer    tracing.AppTracer
}

func NewGameCreatedConsumer(
	logger logger.Logger,
	validator *validator.Validate,
	tracer tracing.AppTracer,
) consumer.ConsumerHandler {
	return &gameCreatedConsumer{logger: logger, validator: validator, tracer: tracer}
}

func (c *gameCreatedConsumer) Handle(
	ctx context.Context,
	consumeContext types.MessageConsumeContext,
) error {
	game, ok := consumeContext.Message().(*GameCreatedV1)
	if !ok {
		return errors.New("error in casting message to GameCreatedV1")
	}

	ctx, span := c.tracer.Start(ctx, "gameCreatedConsumer.Handle")
	span.SetAttributes(attribute.Object("Message", consumeContext.Message()))
	defer span.End()

	command, err := commands.NewCreateGame(
		game.GameId,
		game.Name,
		game.Description,
		game.Price,
		game.Genre,
		game.CreatedAt,
	)
	if err != nil {
		validationErr := customErrors.NewValidationErrorWrap(
			err,
			"[gameCreatedConsumer_Handle.StructCtx] command validation failed",
		)
		c.logger.Errorf(
			fmt.Sprintf(
				"[gameCreatedConsumer_Handle.StructCtx] err: {%v}",
				messageTracing.TraceMessagingErrFromSpan(span, validationErr),
			),
		)

		return err
	}
	_, err = mediatr.Send[*commands.CreateGame, *dtos.CreateGameResponseDto](ctx, command)
	if err != nil {
		err = errors.WithMessage(
			err,
			"[gameCreatedConsumer_Handle.Send] error in sending CreateGame",
		)
		c.logger.Errorw(
			fmt.Sprintf(
				"[gameCreatedConsumer_Handle.Send] id: {%s}, err: {%v}",
				command.GameId,
				messageTracing.TraceMessagingErrFromSpan(span, err),
			),
			logger.Fields{"Id": command.GameId},
		)
	}
	c.logger.Info("Game consumer handled.")

	return err
}
