// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	appendResult "github.com/zizouhuweidi/retro-station/internal/pkg/es/models/append_result"

	expectedStreamVersion "github.com/zizouhuweidi/retro-station/internal/pkg/es/models/stream_version"

	mock "github.com/stretchr/testify/mock"

	models "github.com/zizouhuweidi/retro-station/internal/pkg/es/models"

	readPosition "github.com/zizouhuweidi/retro-station/internal/pkg/es/models/stream_position/read_position"

	streamName "github.com/zizouhuweidi/retro-station/internal/pkg/es/models/stream_name"

	truncatePosition "github.com/zizouhuweidi/retro-station/internal/pkg/es/models/stream_position/truncatePosition"
)

// EventStore is an autogenerated mock type for the EventStore type
type EventStore struct {
	mock.Mock
}

// AppendEvents provides a mock function with given fields: _a0, expectedVersion, events, ctx
func (_m *EventStore) AppendEvents(_a0 streamName.StreamName, expectedVersion expectedStreamVersion.ExpectedStreamVersion, events []*models.StreamEvent, ctx context.Context) (*appendResult.AppendEventsResult, error) {
	ret := _m.Called(_a0, expectedVersion, events, ctx)

	if len(ret) == 0 {
		panic("no return value specified for AppendEvents")
	}

	var r0 *appendResult.AppendEventsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, expectedStreamVersion.ExpectedStreamVersion, []*models.StreamEvent, context.Context) (*appendResult.AppendEventsResult, error)); ok {
		return rf(_a0, expectedVersion, events, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, expectedStreamVersion.ExpectedStreamVersion, []*models.StreamEvent, context.Context) *appendResult.AppendEventsResult); ok {
		r0 = rf(_a0, expectedVersion, events, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appendResult.AppendEventsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, expectedStreamVersion.ExpectedStreamVersion, []*models.StreamEvent, context.Context) error); ok {
		r1 = rf(_a0, expectedVersion, events, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppendNewEvents provides a mock function with given fields: _a0, events, ctx
func (_m *EventStore) AppendNewEvents(_a0 streamName.StreamName, events []*models.StreamEvent, ctx context.Context) (*appendResult.AppendEventsResult, error) {
	ret := _m.Called(_a0, events, ctx)

	if len(ret) == 0 {
		panic("no return value specified for AppendNewEvents")
	}

	var r0 *appendResult.AppendEventsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, []*models.StreamEvent, context.Context) (*appendResult.AppendEventsResult, error)); ok {
		return rf(_a0, events, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, []*models.StreamEvent, context.Context) *appendResult.AppendEventsResult); ok {
		r0 = rf(_a0, events, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appendResult.AppendEventsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, []*models.StreamEvent, context.Context) error); ok {
		r1 = rf(_a0, events, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteStream provides a mock function with given fields: _a0, expectedVersion, ctx
func (_m *EventStore) DeleteStream(_a0 streamName.StreamName, expectedVersion expectedStreamVersion.ExpectedStreamVersion, ctx context.Context) error {
	ret := _m.Called(_a0, expectedVersion, ctx)

	if len(ret) == 0 {
		panic("no return value specified for DeleteStream")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, expectedStreamVersion.ExpectedStreamVersion, context.Context) error); ok {
		r0 = rf(_a0, expectedVersion, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadEvents provides a mock function with given fields: _a0, _a1, count, ctx
func (_m *EventStore) ReadEvents(_a0 streamName.StreamName, _a1 readPosition.StreamReadPosition, count uint64, ctx context.Context) ([]*models.StreamEvent, error) {
	ret := _m.Called(_a0, _a1, count, ctx)

	if len(ret) == 0 {
		panic("no return value specified for ReadEvents")
	}

	var r0 []*models.StreamEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, readPosition.StreamReadPosition, uint64, context.Context) ([]*models.StreamEvent, error)); ok {
		return rf(_a0, _a1, count, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, readPosition.StreamReadPosition, uint64, context.Context) []*models.StreamEvent); ok {
		r0 = rf(_a0, _a1, count, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.StreamEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, readPosition.StreamReadPosition, uint64, context.Context) error); ok {
		r1 = rf(_a0, _a1, count, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadEventsBackwards provides a mock function with given fields: _a0, _a1, count, ctx
func (_m *EventStore) ReadEventsBackwards(_a0 streamName.StreamName, _a1 readPosition.StreamReadPosition, count uint64, ctx context.Context) ([]*models.StreamEvent, error) {
	ret := _m.Called(_a0, _a1, count, ctx)

	if len(ret) == 0 {
		panic("no return value specified for ReadEventsBackwards")
	}

	var r0 []*models.StreamEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, readPosition.StreamReadPosition, uint64, context.Context) ([]*models.StreamEvent, error)); ok {
		return rf(_a0, _a1, count, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, readPosition.StreamReadPosition, uint64, context.Context) []*models.StreamEvent); ok {
		r0 = rf(_a0, _a1, count, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.StreamEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, readPosition.StreamReadPosition, uint64, context.Context) error); ok {
		r1 = rf(_a0, _a1, count, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadEventsBackwardsFromEnd provides a mock function with given fields: _a0, count, ctx
func (_m *EventStore) ReadEventsBackwardsFromEnd(_a0 streamName.StreamName, count uint64, ctx context.Context) ([]*models.StreamEvent, error) {
	ret := _m.Called(_a0, count, ctx)

	if len(ret) == 0 {
		panic("no return value specified for ReadEventsBackwardsFromEnd")
	}

	var r0 []*models.StreamEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, uint64, context.Context) ([]*models.StreamEvent, error)); ok {
		return rf(_a0, count, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, uint64, context.Context) []*models.StreamEvent); ok {
		r0 = rf(_a0, count, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.StreamEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, uint64, context.Context) error); ok {
		r1 = rf(_a0, count, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadEventsBackwardsWithMaxCount provides a mock function with given fields: stream, _a1, ctx
func (_m *EventStore) ReadEventsBackwardsWithMaxCount(stream streamName.StreamName, _a1 readPosition.StreamReadPosition, ctx context.Context) ([]*models.StreamEvent, error) {
	ret := _m.Called(stream, _a1, ctx)

	if len(ret) == 0 {
		panic("no return value specified for ReadEventsBackwardsWithMaxCount")
	}

	var r0 []*models.StreamEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, readPosition.StreamReadPosition, context.Context) ([]*models.StreamEvent, error)); ok {
		return rf(stream, _a1, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, readPosition.StreamReadPosition, context.Context) []*models.StreamEvent); ok {
		r0 = rf(stream, _a1, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.StreamEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, readPosition.StreamReadPosition, context.Context) error); ok {
		r1 = rf(stream, _a1, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadEventsFromStart provides a mock function with given fields: _a0, count, ctx
func (_m *EventStore) ReadEventsFromStart(_a0 streamName.StreamName, count uint64, ctx context.Context) ([]*models.StreamEvent, error) {
	ret := _m.Called(_a0, count, ctx)

	if len(ret) == 0 {
		panic("no return value specified for ReadEventsFromStart")
	}

	var r0 []*models.StreamEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, uint64, context.Context) ([]*models.StreamEvent, error)); ok {
		return rf(_a0, count, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, uint64, context.Context) []*models.StreamEvent); ok {
		r0 = rf(_a0, count, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.StreamEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, uint64, context.Context) error); ok {
		r1 = rf(_a0, count, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadEventsWithMaxCount provides a mock function with given fields: _a0, _a1, ctx
func (_m *EventStore) ReadEventsWithMaxCount(_a0 streamName.StreamName, _a1 readPosition.StreamReadPosition, ctx context.Context) ([]*models.StreamEvent, error) {
	ret := _m.Called(_a0, _a1, ctx)

	if len(ret) == 0 {
		panic("no return value specified for ReadEventsWithMaxCount")
	}

	var r0 []*models.StreamEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, readPosition.StreamReadPosition, context.Context) ([]*models.StreamEvent, error)); ok {
		return rf(_a0, _a1, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, readPosition.StreamReadPosition, context.Context) []*models.StreamEvent); ok {
		r0 = rf(_a0, _a1, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.StreamEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, readPosition.StreamReadPosition, context.Context) error); ok {
		r1 = rf(_a0, _a1, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamExists provides a mock function with given fields: _a0, ctx
func (_m *EventStore) StreamExists(_a0 streamName.StreamName, ctx context.Context) (bool, error) {
	ret := _m.Called(_a0, ctx)

	if len(ret) == 0 {
		panic("no return value specified for StreamExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, context.Context) (bool, error)); ok {
		return rf(_a0, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, context.Context) bool); ok {
		r0 = rf(_a0, ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, context.Context) error); ok {
		r1 = rf(_a0, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TruncateStream provides a mock function with given fields: _a0, _a1, expectedVersion, ctx
func (_m *EventStore) TruncateStream(_a0 streamName.StreamName, _a1 truncatePosition.StreamTruncatePosition, expectedVersion expectedStreamVersion.ExpectedStreamVersion, ctx context.Context) (*appendResult.AppendEventsResult, error) {
	ret := _m.Called(_a0, _a1, expectedVersion, ctx)

	if len(ret) == 0 {
		panic("no return value specified for TruncateStream")
	}

	var r0 *appendResult.AppendEventsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(streamName.StreamName, truncatePosition.StreamTruncatePosition, expectedStreamVersion.ExpectedStreamVersion, context.Context) (*appendResult.AppendEventsResult, error)); ok {
		return rf(_a0, _a1, expectedVersion, ctx)
	}
	if rf, ok := ret.Get(0).(func(streamName.StreamName, truncatePosition.StreamTruncatePosition, expectedStreamVersion.ExpectedStreamVersion, context.Context) *appendResult.AppendEventsResult); ok {
		r0 = rf(_a0, _a1, expectedVersion, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appendResult.AppendEventsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(streamName.StreamName, truncatePosition.StreamTruncatePosition, expectedStreamVersion.ExpectedStreamVersion, context.Context) error); ok {
		r1 = rf(_a0, _a1, expectedVersion, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEventStore creates a new instance of EventStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventStore {
	mock := &EventStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
