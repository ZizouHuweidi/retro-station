// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	consumer "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/consumer"

	metadata "github.com/zizouhuweidi/retro-station/internal/pkg/core/metadata"

	mock "github.com/stretchr/testify/mock"

	types "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
)

// Bus is an autogenerated mock type for the Bus type
type Bus struct {
	mock.Mock
}

// ConnectConsumer provides a mock function with given fields: messageType, _a1
func (_m *Bus) ConnectConsumer(messageType types.IMessage, _a1 consumer.Consumer) error {
	ret := _m.Called(messageType, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ConnectConsumer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.IMessage, consumer.Consumer) error); ok {
		r0 = rf(messageType, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectConsumerHandler provides a mock function with given fields: messageType, consumerHandler
func (_m *Bus) ConnectConsumerHandler(messageType types.IMessage, consumerHandler consumer.ConsumerHandler) error {
	ret := _m.Called(messageType, consumerHandler)

	if len(ret) == 0 {
		panic("no return value specified for ConnectConsumerHandler")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.IMessage, consumer.ConsumerHandler) error); ok {
		r0 = rf(messageType, consumerHandler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsConsumed provides a mock function with given fields: _a0
func (_m *Bus) IsConsumed(_a0 func(types.IMessage)) {
	_m.Called(_a0)
}

// IsProduced provides a mock function with given fields: _a0
func (_m *Bus) IsProduced(_a0 func(types.IMessage)) {
	_m.Called(_a0)
}

// PublishMessage provides a mock function with given fields: ctx, message, meta
func (_m *Bus) PublishMessage(ctx context.Context, message types.IMessage, meta metadata.Metadata) error {
	ret := _m.Called(ctx, message, meta)

	if len(ret) == 0 {
		panic("no return value specified for PublishMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IMessage, metadata.Metadata) error); ok {
		r0 = rf(ctx, message, meta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishMessageWithTopicName provides a mock function with given fields: ctx, message, meta, topicOrExchangeName
func (_m *Bus) PublishMessageWithTopicName(ctx context.Context, message types.IMessage, meta metadata.Metadata, topicOrExchangeName string) error {
	ret := _m.Called(ctx, message, meta, topicOrExchangeName)

	if len(ret) == 0 {
		panic("no return value specified for PublishMessageWithTopicName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IMessage, metadata.Metadata, string) error); ok {
		r0 = rf(ctx, message, meta, topicOrExchangeName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: ctx
func (_m *Bus) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Bus) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBus creates a new instance of Bus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBus(t interface {
	mock.TestingT
	Cleanup(func())
}) *Bus {
	mock := &Bus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
