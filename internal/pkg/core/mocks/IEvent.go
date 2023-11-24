// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/satori/go.uuid"
)

// IEvent is an autogenerated mock type for the IEvent type
type IEvent struct {
	mock.Mock
}

// GetEventId provides a mock function with given fields:
func (_m *IEvent) GetEventId() uuid.UUID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEventId")
	}

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// GetEventType provides a mock function with given fields:
func (_m *IEvent) GetEventType() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEventType")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetOccurredOn provides a mock function with given fields:
func (_m *IEvent) GetOccurredOn() time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOccurredOn")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// NewIEvent creates a new instance of IEvent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIEvent(t interface {
	mock.TestingT
	Cleanup(func())
}) *IEvent {
	mock := &IEvent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
