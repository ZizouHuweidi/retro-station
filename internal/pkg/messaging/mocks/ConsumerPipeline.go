// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	context "context"

	pipeline "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/pipeline"
	types "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	mock "github.com/stretchr/testify/mock"
)

// ConsumerPipeline is an autogenerated mock type for the ConsumerPipeline type
type ConsumerPipeline struct {
	mock.Mock
}

type ConsumerPipeline_Expecter struct {
	mock *mock.Mock
}

func (_m *ConsumerPipeline) EXPECT() *ConsumerPipeline_Expecter {
	return &ConsumerPipeline_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: ctx, consumerContext, next
func (_m *ConsumerPipeline) Handle(ctx context.Context, consumerContext types.MessageConsumeContext, next pipeline.ConsumerHandlerFunc) error {
	ret := _m.Called(ctx, consumerContext, next)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.MessageConsumeContext, pipeline.ConsumerHandlerFunc) error); ok {
		r0 = rf(ctx, consumerContext, next)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConsumerPipeline_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type ConsumerPipeline_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - ctx context.Context
//   - consumerContext types.MessageConsumeContext
//   - next pipeline.ConsumerHandlerFunc
func (_e *ConsumerPipeline_Expecter) Handle(ctx interface{}, consumerContext interface{}, next interface{}) *ConsumerPipeline_Handle_Call {
	return &ConsumerPipeline_Handle_Call{Call: _e.mock.On("Handle", ctx, consumerContext, next)}
}

func (_c *ConsumerPipeline_Handle_Call) Run(run func(ctx context.Context, consumerContext types.MessageConsumeContext, next pipeline.ConsumerHandlerFunc)) *ConsumerPipeline_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.MessageConsumeContext), args[2].(pipeline.ConsumerHandlerFunc))
	})
	return _c
}

func (_c *ConsumerPipeline_Handle_Call) Return(_a0 error) *ConsumerPipeline_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConsumerPipeline_Handle_Call) RunAndReturn(run func(context.Context, types.MessageConsumeContext, pipeline.ConsumerHandlerFunc) error) *ConsumerPipeline_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewConsumerPipeline creates a new instance of ConsumerPipeline. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsumerPipeline(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConsumerPipeline {
	mock := &ConsumerPipeline{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
