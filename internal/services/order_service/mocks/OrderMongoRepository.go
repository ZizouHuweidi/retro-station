// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	context "context"

	read_models "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orderservice/internal/orders/models/orders/read_models"
	mock "github.com/stretchr/testify/mock"

	utils "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/utils"

	uuid "github.com/satori/go.uuid"
)

// OrderMongoRepository is an autogenerated mock type for the OrderMongoRepository type
type OrderMongoRepository struct {
	mock.Mock
}

type OrderMongoRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderMongoRepository) EXPECT() *OrderMongoRepository_Expecter {
	return &OrderMongoRepository_Expecter{mock: &_m.Mock}
}

// CreateOrder provides a mock function with given fields: ctx, order
func (_m *OrderMongoRepository) CreateOrder(ctx context.Context, order *read_models.OrderReadModel) (*read_models.OrderReadModel, error) {
	ret := _m.Called(ctx, order)

	var r0 *read_models.OrderReadModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *read_models.OrderReadModel) (*read_models.OrderReadModel, error)); ok {
		return rf(ctx, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *read_models.OrderReadModel) *read_models.OrderReadModel); ok {
		r0 = rf(ctx, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*read_models.OrderReadModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *read_models.OrderReadModel) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderMongoRepository_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type OrderMongoRepository_CreateOrder_Call struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - order *read_models.OrderReadModel
func (_e *OrderMongoRepository_Expecter) CreateOrder(ctx interface{}, order interface{}) *OrderMongoRepository_CreateOrder_Call {
	return &OrderMongoRepository_CreateOrder_Call{Call: _e.mock.On("CreateOrder", ctx, order)}
}

func (_c *OrderMongoRepository_CreateOrder_Call) Run(run func(ctx context.Context, order *read_models.OrderReadModel)) *OrderMongoRepository_CreateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*read_models.OrderReadModel))
	})
	return _c
}

func (_c *OrderMongoRepository_CreateOrder_Call) Return(_a0 *read_models.OrderReadModel, _a1 error) *OrderMongoRepository_CreateOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderMongoRepository_CreateOrder_Call) RunAndReturn(run func(context.Context, *read_models.OrderReadModel) (*read_models.OrderReadModel, error)) *OrderMongoRepository_CreateOrder_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteOrderByID provides a mock function with given fields: ctx, _a1
func (_m *OrderMongoRepository) DeleteOrderByID(ctx context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderMongoRepository_DeleteOrderByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteOrderByID'
type OrderMongoRepository_DeleteOrderByID_Call struct {
	*mock.Call
}

// DeleteOrderByID is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 uuid.UUID
func (_e *OrderMongoRepository_Expecter) DeleteOrderByID(ctx interface{}, _a1 interface{}) *OrderMongoRepository_DeleteOrderByID_Call {
	return &OrderMongoRepository_DeleteOrderByID_Call{Call: _e.mock.On("DeleteOrderByID", ctx, _a1)}
}

func (_c *OrderMongoRepository_DeleteOrderByID_Call) Run(run func(ctx context.Context, _a1 uuid.UUID)) *OrderMongoRepository_DeleteOrderByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *OrderMongoRepository_DeleteOrderByID_Call) Return(_a0 error) *OrderMongoRepository_DeleteOrderByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrderMongoRepository_DeleteOrderByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *OrderMongoRepository_DeleteOrderByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllOrders provides a mock function with given fields: ctx, listQuery
func (_m *OrderMongoRepository) GetAllOrders(ctx context.Context, listQuery *utils.ListQuery) (*utils.ListResult[*read_models.OrderReadModel], error) {
	ret := _m.Called(ctx, listQuery)

	var r0 *utils.ListResult[*read_models.OrderReadModel]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *utils.ListQuery) (*utils.ListResult[*read_models.OrderReadModel], error)); ok {
		return rf(ctx, listQuery)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *utils.ListQuery) *utils.ListResult[*read_models.OrderReadModel]); ok {
		r0 = rf(ctx, listQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ListResult[*read_models.OrderReadModel])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *utils.ListQuery) error); ok {
		r1 = rf(ctx, listQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderMongoRepository_GetAllOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllOrders'
type OrderMongoRepository_GetAllOrders_Call struct {
	*mock.Call
}

// GetAllOrders is a helper method to define mock.On call
//   - ctx context.Context
//   - listQuery *utils.ListQuery
func (_e *OrderMongoRepository_Expecter) GetAllOrders(ctx interface{}, listQuery interface{}) *OrderMongoRepository_GetAllOrders_Call {
	return &OrderMongoRepository_GetAllOrders_Call{Call: _e.mock.On("GetAllOrders", ctx, listQuery)}
}

func (_c *OrderMongoRepository_GetAllOrders_Call) Run(run func(ctx context.Context, listQuery *utils.ListQuery)) *OrderMongoRepository_GetAllOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*utils.ListQuery))
	})
	return _c
}

func (_c *OrderMongoRepository_GetAllOrders_Call) Return(_a0 *utils.ListResult[*read_models.OrderReadModel], _a1 error) *OrderMongoRepository_GetAllOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderMongoRepository_GetAllOrders_Call) RunAndReturn(run func(context.Context, *utils.ListQuery) (*utils.ListResult[*read_models.OrderReadModel], error)) *OrderMongoRepository_GetAllOrders_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrderById provides a mock function with given fields: ctx, _a1
func (_m *OrderMongoRepository) GetOrderById(ctx context.Context, _a1 uuid.UUID) (*read_models.OrderReadModel, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *read_models.OrderReadModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*read_models.OrderReadModel, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *read_models.OrderReadModel); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*read_models.OrderReadModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderMongoRepository_GetOrderById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderById'
type OrderMongoRepository_GetOrderById_Call struct {
	*mock.Call
}

// GetOrderById is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 uuid.UUID
func (_e *OrderMongoRepository_Expecter) GetOrderById(ctx interface{}, _a1 interface{}) *OrderMongoRepository_GetOrderById_Call {
	return &OrderMongoRepository_GetOrderById_Call{Call: _e.mock.On("GetOrderById", ctx, _a1)}
}

func (_c *OrderMongoRepository_GetOrderById_Call) Run(run func(ctx context.Context, _a1 uuid.UUID)) *OrderMongoRepository_GetOrderById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *OrderMongoRepository_GetOrderById_Call) Return(_a0 *read_models.OrderReadModel, _a1 error) *OrderMongoRepository_GetOrderById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderMongoRepository_GetOrderById_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*read_models.OrderReadModel, error)) *OrderMongoRepository_GetOrderById_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrderByOrderId provides a mock function with given fields: ctx, orderId
func (_m *OrderMongoRepository) GetOrderByOrderId(ctx context.Context, orderId uuid.UUID) (*read_models.OrderReadModel, error) {
	ret := _m.Called(ctx, orderId)

	var r0 *read_models.OrderReadModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*read_models.OrderReadModel, error)); ok {
		return rf(ctx, orderId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *read_models.OrderReadModel); ok {
		r0 = rf(ctx, orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*read_models.OrderReadModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderMongoRepository_GetOrderByOrderId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderByOrderId'
type OrderMongoRepository_GetOrderByOrderId_Call struct {
	*mock.Call
}

// GetOrderByOrderId is a helper method to define mock.On call
//   - ctx context.Context
//   - orderId uuid.UUID
func (_e *OrderMongoRepository_Expecter) GetOrderByOrderId(ctx interface{}, orderId interface{}) *OrderMongoRepository_GetOrderByOrderId_Call {
	return &OrderMongoRepository_GetOrderByOrderId_Call{Call: _e.mock.On("GetOrderByOrderId", ctx, orderId)}
}

func (_c *OrderMongoRepository_GetOrderByOrderId_Call) Run(run func(ctx context.Context, orderId uuid.UUID)) *OrderMongoRepository_GetOrderByOrderId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *OrderMongoRepository_GetOrderByOrderId_Call) Return(_a0 *read_models.OrderReadModel, _a1 error) *OrderMongoRepository_GetOrderByOrderId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderMongoRepository_GetOrderByOrderId_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*read_models.OrderReadModel, error)) *OrderMongoRepository_GetOrderByOrderId_Call {
	_c.Call.Return(run)
	return _c
}

// SearchOrders provides a mock function with given fields: ctx, searchText, listQuery
func (_m *OrderMongoRepository) SearchOrders(ctx context.Context, searchText string, listQuery *utils.ListQuery) (*utils.ListResult[*read_models.OrderReadModel], error) {
	ret := _m.Called(ctx, searchText, listQuery)

	var r0 *utils.ListResult[*read_models.OrderReadModel]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *utils.ListQuery) (*utils.ListResult[*read_models.OrderReadModel], error)); ok {
		return rf(ctx, searchText, listQuery)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *utils.ListQuery) *utils.ListResult[*read_models.OrderReadModel]); ok {
		r0 = rf(ctx, searchText, listQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ListResult[*read_models.OrderReadModel])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *utils.ListQuery) error); ok {
		r1 = rf(ctx, searchText, listQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderMongoRepository_SearchOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchOrders'
type OrderMongoRepository_SearchOrders_Call struct {
	*mock.Call
}

// SearchOrders is a helper method to define mock.On call
//   - ctx context.Context
//   - searchText string
//   - listQuery *utils.ListQuery
func (_e *OrderMongoRepository_Expecter) SearchOrders(ctx interface{}, searchText interface{}, listQuery interface{}) *OrderMongoRepository_SearchOrders_Call {
	return &OrderMongoRepository_SearchOrders_Call{Call: _e.mock.On("SearchOrders", ctx, searchText, listQuery)}
}

func (_c *OrderMongoRepository_SearchOrders_Call) Run(run func(ctx context.Context, searchText string, listQuery *utils.ListQuery)) *OrderMongoRepository_SearchOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*utils.ListQuery))
	})
	return _c
}

func (_c *OrderMongoRepository_SearchOrders_Call) Return(_a0 *utils.ListResult[*read_models.OrderReadModel], _a1 error) *OrderMongoRepository_SearchOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderMongoRepository_SearchOrders_Call) RunAndReturn(run func(context.Context, string, *utils.ListQuery) (*utils.ListResult[*read_models.OrderReadModel], error)) *OrderMongoRepository_SearchOrders_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOrder provides a mock function with given fields: ctx, order
func (_m *OrderMongoRepository) UpdateOrder(ctx context.Context, order *read_models.OrderReadModel) (*read_models.OrderReadModel, error) {
	ret := _m.Called(ctx, order)

	var r0 *read_models.OrderReadModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *read_models.OrderReadModel) (*read_models.OrderReadModel, error)); ok {
		return rf(ctx, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *read_models.OrderReadModel) *read_models.OrderReadModel); ok {
		r0 = rf(ctx, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*read_models.OrderReadModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *read_models.OrderReadModel) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderMongoRepository_UpdateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOrder'
type OrderMongoRepository_UpdateOrder_Call struct {
	*mock.Call
}

// UpdateOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - order *read_models.OrderReadModel
func (_e *OrderMongoRepository_Expecter) UpdateOrder(ctx interface{}, order interface{}) *OrderMongoRepository_UpdateOrder_Call {
	return &OrderMongoRepository_UpdateOrder_Call{Call: _e.mock.On("UpdateOrder", ctx, order)}
}

func (_c *OrderMongoRepository_UpdateOrder_Call) Run(run func(ctx context.Context, order *read_models.OrderReadModel)) *OrderMongoRepository_UpdateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*read_models.OrderReadModel))
	})
	return _c
}

func (_c *OrderMongoRepository_UpdateOrder_Call) Return(_a0 *read_models.OrderReadModel, _a1 error) *OrderMongoRepository_UpdateOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderMongoRepository_UpdateOrder_Call) RunAndReturn(run func(context.Context, *read_models.OrderReadModel) (*read_models.OrderReadModel, error)) *OrderMongoRepository_UpdateOrder_Call {
	_c.Call.Return(run)
	return _c
}

// NewOrderMongoRepository creates a new instance of OrderMongoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderMongoRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderMongoRepository {
	mock := &OrderMongoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
