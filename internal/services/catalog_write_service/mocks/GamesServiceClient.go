// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	games_service "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc/genproto"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// GamesServiceClient is an autogenerated mock type for the GamesServiceClient type
type GamesServiceClient struct {
	mock.Mock
}

// CreateGame provides a mock function with given fields: ctx, in, opts
func (_m *GamesServiceClient) CreateGame(ctx context.Context, in *games_service.CreateGameReq, opts ...grpc.CallOption) (*games_service.CreateGameRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateGame")
	}

	var r0 *games_service.CreateGameRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *games_service.CreateGameReq, ...grpc.CallOption) (*games_service.CreateGameRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *games_service.CreateGameReq, ...grpc.CallOption) *games_service.CreateGameRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*games_service.CreateGameRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *games_service.CreateGameReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGameById provides a mock function with given fields: ctx, in, opts
func (_m *GamesServiceClient) GetGameById(ctx context.Context, in *games_service.GetGameByIdReq, opts ...grpc.CallOption) (*games_service.GetGameByIdRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGameById")
	}

	var r0 *games_service.GetGameByIdRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *games_service.GetGameByIdReq, ...grpc.CallOption) (*games_service.GetGameByIdRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *games_service.GetGameByIdReq, ...grpc.CallOption) *games_service.GetGameByIdRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*games_service.GetGameByIdRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *games_service.GetGameByIdReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGame provides a mock function with given fields: ctx, in, opts
func (_m *GamesServiceClient) UpdateGame(ctx context.Context, in *games_service.UpdateGameReq, opts ...grpc.CallOption) (*games_service.UpdateGameRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGame")
	}

	var r0 *games_service.UpdateGameRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *games_service.UpdateGameReq, ...grpc.CallOption) (*games_service.UpdateGameRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *games_service.UpdateGameReq, ...grpc.CallOption) *games_service.UpdateGameRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*games_service.UpdateGameRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *games_service.UpdateGameReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGamesServiceClient creates a new instance of GamesServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGamesServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *GamesServiceClient {
	mock := &GamesServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
