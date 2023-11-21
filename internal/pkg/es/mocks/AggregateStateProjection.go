// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	domain "github.com/zizouhuweidi/retro-station/internal/pkg/core/domain"
	metadata "github.com/zizouhuweidi/retro-station/internal/pkg/core/metadata"

	mock "github.com/stretchr/testify/mock"
)

// AggregateStateProjection is an autogenerated mock type for the AggregateStateProjection type
type AggregateStateProjection struct {
	mock.Mock
}

type AggregateStateProjection_Expecter struct {
	mock *mock.Mock
}

func (_m *AggregateStateProjection) EXPECT() *AggregateStateProjection_Expecter {
	return &AggregateStateProjection_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: event, isNew
func (_m *AggregateStateProjection) Apply(event domain.IDomainEvent, isNew bool) error {
	ret := _m.Called(event, isNew)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.IDomainEvent, bool) error); ok {
		r0 = rf(event, isNew)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AggregateStateProjection_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type AggregateStateProjection_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - event domain.IDomainEvent
//   - isNew bool
func (_e *AggregateStateProjection_Expecter) Apply(event interface{}, isNew interface{}) *AggregateStateProjection_Apply_Call {
	return &AggregateStateProjection_Apply_Call{Call: _e.mock.On("Apply", event, isNew)}
}

func (_c *AggregateStateProjection_Apply_Call) Run(run func(event domain.IDomainEvent, isNew bool)) *AggregateStateProjection_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.IDomainEvent), args[1].(bool))
	})
	return _c
}

func (_c *AggregateStateProjection_Apply_Call) Return(_a0 error) *AggregateStateProjection_Apply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AggregateStateProjection_Apply_Call) RunAndReturn(run func(domain.IDomainEvent, bool) error) *AggregateStateProjection_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// fold provides a mock function with given fields: event, _a1
func (_m *AggregateStateProjection) fold(event domain.IDomainEvent, _a1 metadata.Metadata) error {
	ret := _m.Called(event, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.IDomainEvent, metadata.Metadata) error); ok {
		r0 = rf(event, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AggregateStateProjection_fold_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'fold'
type AggregateStateProjection_fold_Call struct {
	*mock.Call
}

// fold is a helper method to define mock.On call
//   - event domain.IDomainEvent
//   - _a1 metadata.Metadata
func (_e *AggregateStateProjection_Expecter) fold(event interface{}, _a1 interface{}) *AggregateStateProjection_fold_Call {
	return &AggregateStateProjection_fold_Call{Call: _e.mock.On("fold", event, _a1)}
}

func (_c *AggregateStateProjection_fold_Call) Run(run func(event domain.IDomainEvent, _a1 metadata.Metadata)) *AggregateStateProjection_fold_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.IDomainEvent), args[1].(metadata.Metadata))
	})
	return _c
}

func (_c *AggregateStateProjection_fold_Call) Return(_a0 error) *AggregateStateProjection_fold_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AggregateStateProjection_fold_Call) RunAndReturn(run func(domain.IDomainEvent, metadata.Metadata) error) *AggregateStateProjection_fold_Call {
	_c.Call.Return(run)
	return _c
}

// NewAggregateStateProjection creates a new instance of AggregateStateProjection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAggregateStateProjection(t interface {
	mock.TestingT
	Cleanup(func())
}) *AggregateStateProjection {
	mock := &AggregateStateProjection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}