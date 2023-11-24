// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/zizouhuweidi/retro-station/internal/pkg/es/models"
)

// IProjection is an autogenerated mock type for the IProjection type
type IProjection struct {
	mock.Mock
}

// ProcessEvent provides a mock function with given fields: ctx, streamEvent
func (_m *IProjection) ProcessEvent(ctx context.Context, streamEvent *models.StreamEvent) error {
	ret := _m.Called(ctx, streamEvent)

	if len(ret) == 0 {
		panic("no return value specified for ProcessEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.StreamEvent) error); ok {
		r0 = rf(ctx, streamEvent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIProjection creates a new instance of IProjection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProjection(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProjection {
	mock := &IProjection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
