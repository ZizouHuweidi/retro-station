package externalEvents

import (
	"context"
	"fmt"

	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/consumer"
	messageTracing "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/products/features/deleting_products/v1/commands"

	"emperror.dev/errors"
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/go-mediatr"
)

type productDeletedConsumer struct {
	logger    logger.Logger
	validator *validator.Validate
	tracer    tracing.AppTracer
}

func NewProductDeletedConsumer(
	logger logger.Logger,
	validator *validator.Validate,
	tracer tracing.AppTracer,
) consumer.ConsumerHandler {
	return &productDeletedConsumer{logger: logger, validator: validator, tracer: tracer}
}

func (c *productDeletedConsumer) Handle(
	ctx context.Context,
	consumeContext types.MessageConsumeContext,
) error {
	message, ok := consumeContext.Message().(*ProductDeletedV1)
	if !ok {
		return errors.New("error in casting message to ProductDeletedV1")
	}

	ctx, span := c.tracer.Start(ctx, "productDeletedConsumer.Handle")
	span.SetAttributes(attribute.Object("Message", consumeContext.Message()))
	defer span.End()

	productUUID, err := uuid.FromString(message.ProductId)
	if err != nil {
		badRequestErr := customErrors.NewBadRequestErrorWrap(
			err,
			"[productDeletedConsumer_Handle.uuid.FromString] error in the converting uuid",
		)
		c.logger.Errorf(
			fmt.Sprintf(
				"[productDeletedConsumer_Handle.uuid.FromString] err: %v",
				messageTracing.TraceMessagingErrFromSpan(span, badRequestErr),
			),
		)

		return err
	}

	command, err := commands.NewDeleteProduct(productUUID)
	if err != nil {
		validationErr := customErrors.NewValidationErrorWrap(
			err,
			"[productDeletedConsumer_Handle.StructCtx] command validation failed",
		)
		c.logger.Errorf(
			fmt.Sprintf(
				"[productDeletedConsumer_Consume.StructCtx] err: {%v}",
				messageTracing.TraceMessagingErrFromSpan(span, validationErr),
			),
		)

		return err
	}

	_, err = mediatr.Send[*commands.DeleteProduct, *mediatr.Unit](ctx, command)

	return err
}
