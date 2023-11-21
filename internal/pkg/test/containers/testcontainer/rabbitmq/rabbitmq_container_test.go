package rabbitmq

import (
	"context"
	"testing"
	"time"

	"github.com/zizouhuweidi/retro-station/internal/pkg/core/serializer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/core/serializer/json"
	defaultLogger "github.com/zizouhuweidi/retro-station/internal/pkg/logger/default_logger"
	messageConsumer "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/consumer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	rabbitmqConfigurations "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/configurations"
	consumerConfigurations "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/consumer/configurations"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/messaging/consumer"
	testUtils "github.com/zizouhuweidi/retro-station/internal/pkg/test/utils"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func Test_Custom_RabbitMQ_Container(t *testing.T) {
	ctx := context.Background()
	fakeConsumer := consumer.NewRabbitMQFakeTestConsumerHandler[*ProducerConsumerMessage]()
	defaultLogger.SetupDefaultLogger()
	eventSerializer := serializer.NewDefaultEventSerializer(json.NewDefaultSerializer())

	rabbitmq, err := NewRabbitMQTestContainers(
		defaultLogger.Logger,
	).Start(ctx, t, eventSerializer, func(builder rabbitmqConfigurations.RabbitMQConfigurationBuilder) {
		builder.AddConsumer(ProducerConsumerMessage{},
			func(consumerBuilder consumerConfigurations.RabbitMQConsumerConfigurationBuilder) {
				consumerBuilder.WithHandlers(
					func(handlerBuilder messageConsumer.ConsumerHandlerConfigurationBuilder) {
						handlerBuilder.AddHandler(fakeConsumer)
					},
				)
			})
	})

	require.NoError(t, err)
	require.NotNil(t, rabbitmq)

	err = rabbitmq.Start(ctx)
	require.NoError(t, err)

	// wait for consumers ready to consume before publishing messages (for preventing messages lost)
	time.Sleep(time.Second * 1)

	err = rabbitmq.PublishMessage(
		context.Background(),
		&ProducerConsumerMessage{
			Data:    "ssssssssss",
			Message: types.NewMessage(uuid.NewV4().String()),
		},
		nil,
	)
	if err != nil {
		return
	}

	err = testUtils.WaitUntilConditionMet(func() bool {
		return fakeConsumer.IsHandled()
	})

	t.Log("stopping test container")

	if err != nil {
		require.FailNow(t, err.Error())
	}
}

type ProducerConsumerMessage struct {
	*types.Message
	Data string
}
