package messaging

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/bus"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/utils"
	typeMapper "github.com/zizouhuweidi/retro-station/internal/pkg/reflection/type_mappper"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/hypothesis"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/messaging/consumer"
)

func ShouldProduced[T types.IMessage](
	ctx context.Context,
	bus bus.Bus,
	condition func(T) bool,
) hypothesis.Hypothesis[T] {
	hypo := hypothesis.ForT[T](condition)

	bus.IsProduced(func(message types.IMessage) {
		typ := utils.GetMessageBaseReflectType(typeMapper.GenericInstanceByT[T]())
		if utils.GetMessageBaseReflectType(message) == typ {
			m, ok := message.(T)
			if !ok {
				hypo.Test(ctx, *new(T))
			}
			hypo.Test(ctx, m)
		}
	})

	return hypo
}

func ShouldConsume[T types.IMessage](
	ctx context.Context,
	bus bus.Bus,
	condition func(T) bool,
) hypothesis.Hypothesis[T] {
	hypo := hypothesis.ForT[T](condition)

	bus.IsConsumed(func(message types.IMessage) {
		typ := utils.GetMessageBaseReflectType(typeMapper.GenericInstanceByT[T]())
		if utils.GetMessageBaseReflectType(message) == typ {
			m, ok := message.(T)
			if !ok {
				hypo.Test(ctx, *new(T))
			}
			hypo.Test(ctx, m)
		}
	})

	return hypo
}

func ShouldConsumeNewConsumer[T types.IMessage](bus bus.Bus) (hypothesis.Hypothesis[T], error) {
	hypo := hypothesis.ForT[T](nil)
	testConsumer := consumer.NewRabbitMQFakeTestConsumerHandlerWithHypothesis(hypo)
	err := bus.ConnectConsumerHandler(typeMapper.GenericInstanceByT[T](), testConsumer)
	if err != nil {
		return nil, err
	}

	return hypo, nil
}
