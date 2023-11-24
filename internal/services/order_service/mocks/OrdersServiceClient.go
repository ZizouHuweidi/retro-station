// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	orders_service "github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/shared/grpc/genproto"
)

// OrdersServiceClient is an autogenerated mock type for the OrdersServiceClient type
type OrdersServiceClient struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: ctx, in, opts
func (_m *OrdersServiceClient) CreateOrder(ctx context.Context, in *orders_service.CreateOrderReq, opts ...grpc.CallOption) (*orders_service.CreateOrderRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 *orders_service.CreateOrderRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.CreateOrderReq, ...grpc.CallOption) (*orders_service.CreateOrderRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.CreateOrderReq, ...grpc.CallOption) *orders_service.CreateOrderRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.CreateOrderRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.CreateOrderReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByID provides a mock function with given fields: ctx, in, opts
func (_m *OrdersServiceClient) GetOrderByID(ctx context.Context, in *orders_service.GetOrderByIDReq, opts ...grpc.CallOption) (*orders_service.GetOrderByIDRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderByID")
	}

	var r0 *orders_service.GetOrderByIDRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.GetOrderByIDReq, ...grpc.CallOption) (*orders_service.GetOrderByIDRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.GetOrderByIDReq, ...grpc.CallOption) *orders_service.GetOrderByIDRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.GetOrderByIDRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.GetOrderByIDReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrders provides a mock function with given fields: ctx, in, opts
func (_m *OrdersServiceClient) GetOrders(ctx context.Context, in *orders_service.GetOrdersReq, opts ...grpc.CallOption) (*orders_service.GetOrdersRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 *orders_service.GetOrdersRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.GetOrdersReq, ...grpc.CallOption) (*orders_service.GetOrdersRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.GetOrdersReq, ...grpc.CallOption) *orders_service.GetOrdersRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.GetOrdersRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.GetOrdersReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitOrder provides a mock function with given fields: ctx, in, opts
func (_m *OrdersServiceClient) SubmitOrder(ctx context.Context, in *orders_service.SubmitOrderReq, opts ...grpc.CallOption) (*orders_service.SubmitOrderRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SubmitOrder")
	}

	var r0 *orders_service.SubmitOrderRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.SubmitOrderReq, ...grpc.CallOption) (*orders_service.SubmitOrderRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.SubmitOrderReq, ...grpc.CallOption) *orders_service.SubmitOrderRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.SubmitOrderRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.SubmitOrderReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateShoppingCart provides a mock function with given fields: ctx, in, opts
func (_m *OrdersServiceClient) UpdateShoppingCart(ctx context.Context, in *orders_service.UpdateShoppingCartReq, opts ...grpc.CallOption) (*orders_service.UpdateShoppingCartRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateShoppingCart")
	}

	var r0 *orders_service.UpdateShoppingCartRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.UpdateShoppingCartReq, ...grpc.CallOption) (*orders_service.UpdateShoppingCartRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.UpdateShoppingCartReq, ...grpc.CallOption) *orders_service.UpdateShoppingCartRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.UpdateShoppingCartRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.UpdateShoppingCartReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrdersServiceClient creates a new instance of OrdersServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrdersServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrdersServiceClient {
	mock := &OrdersServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
