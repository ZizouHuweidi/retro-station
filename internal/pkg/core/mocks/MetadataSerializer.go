// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	metadata "github.com/zizouhuweidi/retro-station/internal/pkg/core/metadata"
	mock "github.com/stretchr/testify/mock"
)

// MetadataSerializer is an autogenerated mock type for the MetadataSerializer type
type MetadataSerializer struct {
	mock.Mock
}

type MetadataSerializer_Expecter struct {
	mock *mock.Mock
}

func (_m *MetadataSerializer) EXPECT() *MetadataSerializer_Expecter {
	return &MetadataSerializer_Expecter{mock: &_m.Mock}
}

// Deserialize provides a mock function with given fields: bytes
func (_m *MetadataSerializer) Deserialize(bytes []byte) (metadata.Metadata, error) {
	ret := _m.Called(bytes)

	var r0 metadata.Metadata
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (metadata.Metadata, error)); ok {
		return rf(bytes)
	}
	if rf, ok := ret.Get(0).(func([]byte) metadata.Metadata); ok {
		r0 = rf(bytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.Metadata)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(bytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetadataSerializer_Deserialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Deserialize'
type MetadataSerializer_Deserialize_Call struct {
	*mock.Call
}

// Deserialize is a helper method to define mock.On call
//   - bytes []byte
func (_e *MetadataSerializer_Expecter) Deserialize(bytes interface{}) *MetadataSerializer_Deserialize_Call {
	return &MetadataSerializer_Deserialize_Call{Call: _e.mock.On("Deserialize", bytes)}
}

func (_c *MetadataSerializer_Deserialize_Call) Run(run func(bytes []byte)) *MetadataSerializer_Deserialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MetadataSerializer_Deserialize_Call) Return(_a0 metadata.Metadata, _a1 error) *MetadataSerializer_Deserialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetadataSerializer_Deserialize_Call) RunAndReturn(run func([]byte) (metadata.Metadata, error)) *MetadataSerializer_Deserialize_Call {
	_c.Call.Return(run)
	return _c
}

// Serialize provides a mock function with given fields: meta
func (_m *MetadataSerializer) Serialize(meta metadata.Metadata) ([]byte, error) {
	ret := _m.Called(meta)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(metadata.Metadata) ([]byte, error)); ok {
		return rf(meta)
	}
	if rf, ok := ret.Get(0).(func(metadata.Metadata) []byte); ok {
		r0 = rf(meta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(metadata.Metadata) error); ok {
		r1 = rf(meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetadataSerializer_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type MetadataSerializer_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
//   - meta metadata.Metadata
func (_e *MetadataSerializer_Expecter) Serialize(meta interface{}) *MetadataSerializer_Serialize_Call {
	return &MetadataSerializer_Serialize_Call{Call: _e.mock.On("Serialize", meta)}
}

func (_c *MetadataSerializer_Serialize_Call) Run(run func(meta metadata.Metadata)) *MetadataSerializer_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.Metadata))
	})
	return _c
}

func (_c *MetadataSerializer_Serialize_Call) Return(_a0 []byte, _a1 error) *MetadataSerializer_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetadataSerializer_Serialize_Call) RunAndReturn(run func(metadata.Metadata) ([]byte, error)) *MetadataSerializer_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// NewMetadataSerializer creates a new instance of MetadataSerializer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetadataSerializer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MetadataSerializer {
	mock := &MetadataSerializer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
