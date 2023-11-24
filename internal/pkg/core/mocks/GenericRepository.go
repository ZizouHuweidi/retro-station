// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	specification "github.com/zizouhuweidi/retro-station/internal/pkg/core/data/specification"

	utils "github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	uuid "github.com/satori/go.uuid"
)

// GenericRepository is an autogenerated mock type for the GenericRepository type
type GenericRepository[TEntity interface{}] struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, entity
func (_m *GenericRepository[TEntity]) Add(ctx context.Context, entity TEntity) error {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, TEntity) error); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddAll provides a mock function with given fields: ctx, entities
func (_m *GenericRepository[TEntity]) AddAll(ctx context.Context, entities []TEntity) error {
	ret := _m.Called(ctx, entities)

	if len(ret) == 0 {
		panic("no return value specified for AddAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []TEntity) error); ok {
		r0 = rf(ctx, entities)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields: ctx
func (_m *GenericRepository[TEntity]) Count(ctx context.Context) int64 {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *GenericRepository[TEntity]) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, _a1
func (_m *GenericRepository[TEntity]) Find(ctx context.Context, _a1 specification.Specification) ([]TEntity, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 []TEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, specification.Specification) ([]TEntity, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, specification.Specification) []TEntity); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, specification.Specification) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstOrDefault provides a mock function with given fields: ctx, filters
func (_m *GenericRepository[TEntity]) FirstOrDefault(ctx context.Context, filters map[string]interface{}) (TEntity, error) {
	ret := _m.Called(ctx, filters)

	if len(ret) == 0 {
		panic("no return value specified for FirstOrDefault")
	}

	var r0 TEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) (TEntity, error)); ok {
		return rf(ctx, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) TEntity); ok {
		r0 = rf(ctx, filters)
	} else {
		r0 = ret.Get(0).(TEntity)
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx, listQuery
func (_m *GenericRepository[TEntity]) GetAll(ctx context.Context, listQuery *utils.ListQuery) (*utils.ListResult[TEntity], error) {
	ret := _m.Called(ctx, listQuery)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 *utils.ListResult[TEntity]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *utils.ListQuery) (*utils.ListResult[TEntity], error)); ok {
		return rf(ctx, listQuery)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *utils.ListQuery) *utils.ListResult[TEntity]); ok {
		r0 = rf(ctx, listQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ListResult[TEntity])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *utils.ListQuery) error); ok {
		r1 = rf(ctx, listQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByFilter provides a mock function with given fields: ctx, filters
func (_m *GenericRepository[TEntity]) GetByFilter(ctx context.Context, filters map[string]interface{}) ([]TEntity, error) {
	ret := _m.Called(ctx, filters)

	if len(ret) == 0 {
		panic("no return value specified for GetByFilter")
	}

	var r0 []TEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) ([]TEntity, error)); ok {
		return rf(ctx, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) []TEntity); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByFuncFilter provides a mock function with given fields: ctx, filterFunc
func (_m *GenericRepository[TEntity]) GetByFuncFilter(ctx context.Context, filterFunc func(TEntity) bool) ([]TEntity, error) {
	ret := _m.Called(ctx, filterFunc)

	if len(ret) == 0 {
		panic("no return value specified for GetByFuncFilter")
	}

	var r0 []TEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, func(TEntity) bool) ([]TEntity, error)); ok {
		return rf(ctx, filterFunc)
	}
	if rf, ok := ret.Get(0).(func(context.Context, func(TEntity) bool) []TEntity); ok {
		r0 = rf(ctx, filterFunc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, func(TEntity) bool) error); ok {
		r1 = rf(ctx, filterFunc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *GenericRepository[TEntity]) GetById(ctx context.Context, id uuid.UUID) (TEntity, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 TEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (TEntity, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) TEntity); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(TEntity)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, searchTerm, listQuery
func (_m *GenericRepository[TEntity]) Search(ctx context.Context, searchTerm string, listQuery *utils.ListQuery) (*utils.ListResult[TEntity], error) {
	ret := _m.Called(ctx, searchTerm, listQuery)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 *utils.ListResult[TEntity]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *utils.ListQuery) (*utils.ListResult[TEntity], error)); ok {
		return rf(ctx, searchTerm, listQuery)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *utils.ListQuery) *utils.ListResult[TEntity]); ok {
		r0 = rf(ctx, searchTerm, listQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ListResult[TEntity])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *utils.ListQuery) error); ok {
		r1 = rf(ctx, searchTerm, listQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SkipTake provides a mock function with given fields: ctx, skip, take
func (_m *GenericRepository[TEntity]) SkipTake(ctx context.Context, skip int, take int) ([]TEntity, error) {
	ret := _m.Called(ctx, skip, take)

	if len(ret) == 0 {
		panic("no return value specified for SkipTake")
	}

	var r0 []TEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]TEntity, error)); ok {
		return rf(ctx, skip, take)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []TEntity); ok {
		r0 = rf(ctx, skip, take)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, skip, take)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, entity
func (_m *GenericRepository[TEntity]) Update(ctx context.Context, entity TEntity) error {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, TEntity) error); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAll provides a mock function with given fields: ctx, entities
func (_m *GenericRepository[TEntity]) UpdateAll(ctx context.Context, entities []TEntity) error {
	ret := _m.Called(ctx, entities)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []TEntity) error); ok {
		r0 = rf(ctx, entities)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewGenericRepository creates a new instance of GenericRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGenericRepository[TEntity interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *GenericRepository[TEntity] {
	mock := &GenericRepository[TEntity]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
