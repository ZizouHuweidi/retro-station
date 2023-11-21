// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/zizouhuweidi/retro-station/internal/pkg/es/models"
	mock "github.com/stretchr/testify/mock"
)

// IProjectionPublisher is an autogenerated mock type for the IProjectionPublisher type
type IProjectionPublisher struct {
	mock.Mock
}

type IProjectionPublisher_Expecter struct {
	mock *mock.Mock
}

func (_m *IProjectionPublisher) EXPECT() *IProjectionPublisher_Expecter {
	return &IProjectionPublisher_Expecter{mock: &_m.Mock}
}

// Publish provides a mock function with given fields: ctx, streamEvent
func (_m *IProjectionPublisher) Publish(ctx context.Context, streamEvent *models.StreamEvent) error {
	ret := _m.Called(ctx, streamEvent)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.StreamEvent) error); ok {
		r0 = rf(ctx, streamEvent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IProjectionPublisher_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type IProjectionPublisher_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - ctx context.Context
//   - streamEvent *models.StreamEvent
func (_e *IProjectionPublisher_Expecter) Publish(ctx interface{}, streamEvent interface{}) *IProjectionPublisher_Publish_Call {
	return &IProjectionPublisher_Publish_Call{Call: _e.mock.On("Publish", ctx, streamEvent)}
}

func (_c *IProjectionPublisher_Publish_Call) Run(run func(ctx context.Context, streamEvent *models.StreamEvent)) *IProjectionPublisher_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.StreamEvent))
	})
	return _c
}

func (_c *IProjectionPublisher_Publish_Call) Return(_a0 error) *IProjectionPublisher_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IProjectionPublisher_Publish_Call) RunAndReturn(run func(context.Context, *models.StreamEvent) error) *IProjectionPublisher_Publish_Call {
	_c.Call.Return(run)
	return _c
}

// NewIProjectionPublisher creates a new instance of IProjectionPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProjectionPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProjectionPublisher {
	mock := &IProjectionPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}