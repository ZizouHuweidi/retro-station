// Code generated by mockery v2.30.16. DO NOT EDIT.

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

type GenericRepository_Expecter[TEntity interface{}] struct {
	mock *mock.Mock
}

func (_m *GenericRepository[TEntity]) EXPECT() *GenericRepository_Expecter[TEntity] {
	return &GenericRepository_Expecter[TEntity]{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, entity
func (_m *GenericRepository[TEntity]) Add(ctx context.Context, entity TEntity) error {
	ret := _m.Called(ctx, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, TEntity) error); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenericRepository_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type GenericRepository_Add_Call[TEntity interface{}] struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - entity TEntity
func (_e *GenericRepository_Expecter[TEntity]) Add(ctx interface{}, entity interface{}) *GenericRepository_Add_Call[TEntity] {
	return &GenericRepository_Add_Call[TEntity]{Call: _e.mock.On("Add", ctx, entity)}
}

func (_c *GenericRepository_Add_Call[TEntity]) Run(run func(ctx context.Context, entity TEntity)) *GenericRepository_Add_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(TEntity))
	})
	return _c
}

func (_c *GenericRepository_Add_Call[TEntity]) Return(_a0 error) *GenericRepository_Add_Call[TEntity] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GenericRepository_Add_Call[TEntity]) RunAndReturn(run func(context.Context, TEntity) error) *GenericRepository_Add_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// AddAll provides a mock function with given fields: ctx, entities
func (_m *GenericRepository[TEntity]) AddAll(ctx context.Context, entities []TEntity) error {
	ret := _m.Called(ctx, entities)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []TEntity) error); ok {
		r0 = rf(ctx, entities)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenericRepository_AddAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddAll'
type GenericRepository_AddAll_Call[TEntity interface{}] struct {
	*mock.Call
}

// AddAll is a helper method to define mock.On call
//   - ctx context.Context
//   - entities []TEntity
func (_e *GenericRepository_Expecter[TEntity]) AddAll(ctx interface{}, entities interface{}) *GenericRepository_AddAll_Call[TEntity] {
	return &GenericRepository_AddAll_Call[TEntity]{Call: _e.mock.On("AddAll", ctx, entities)}
}

func (_c *GenericRepository_AddAll_Call[TEntity]) Run(run func(ctx context.Context, entities []TEntity)) *GenericRepository_AddAll_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]TEntity))
	})
	return _c
}

func (_c *GenericRepository_AddAll_Call[TEntity]) Return(_a0 error) *GenericRepository_AddAll_Call[TEntity] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GenericRepository_AddAll_Call[TEntity]) RunAndReturn(run func(context.Context, []TEntity) error) *GenericRepository_AddAll_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// Count provides a mock function with given fields: ctx
func (_m *GenericRepository[TEntity]) Count(ctx context.Context) int64 {
	ret := _m.Called(ctx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GenericRepository_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type GenericRepository_Count_Call[TEntity interface{}] struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
func (_e *GenericRepository_Expecter[TEntity]) Count(ctx interface{}) *GenericRepository_Count_Call[TEntity] {
	return &GenericRepository_Count_Call[TEntity]{Call: _e.mock.On("Count", ctx)}
}

func (_c *GenericRepository_Count_Call[TEntity]) Run(run func(ctx context.Context)) *GenericRepository_Count_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *GenericRepository_Count_Call[TEntity]) Return(_a0 int64) *GenericRepository_Count_Call[TEntity] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GenericRepository_Count_Call[TEntity]) RunAndReturn(run func(context.Context) int64) *GenericRepository_Count_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *GenericRepository[TEntity]) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenericRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type GenericRepository_Delete_Call[TEntity interface{}] struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *GenericRepository_Expecter[TEntity]) Delete(ctx interface{}, id interface{}) *GenericRepository_Delete_Call[TEntity] {
	return &GenericRepository_Delete_Call[TEntity]{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *GenericRepository_Delete_Call[TEntity]) Run(run func(ctx context.Context, id uuid.UUID)) *GenericRepository_Delete_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *GenericRepository_Delete_Call[TEntity]) Return(_a0 error) *GenericRepository_Delete_Call[TEntity] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GenericRepository_Delete_Call[TEntity]) RunAndReturn(run func(context.Context, uuid.UUID) error) *GenericRepository_Delete_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: ctx, _a1
func (_m *GenericRepository[TEntity]) Find(ctx context.Context, _a1 specification.Specification) ([]TEntity, error) {
	ret := _m.Called(ctx, _a1)

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

// GenericRepository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type GenericRepository_Find_Call[TEntity interface{}] struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 specification.Specification
func (_e *GenericRepository_Expecter[TEntity]) Find(ctx interface{}, _a1 interface{}) *GenericRepository_Find_Call[TEntity] {
	return &GenericRepository_Find_Call[TEntity]{Call: _e.mock.On("Find", ctx, _a1)}
}

func (_c *GenericRepository_Find_Call[TEntity]) Run(run func(ctx context.Context, _a1 specification.Specification)) *GenericRepository_Find_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(specification.Specification))
	})
	return _c
}

func (_c *GenericRepository_Find_Call[TEntity]) Return(_a0 []TEntity, _a1 error) *GenericRepository_Find_Call[TEntity] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GenericRepository_Find_Call[TEntity]) RunAndReturn(run func(context.Context, specification.Specification) ([]TEntity, error)) *GenericRepository_Find_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// FirstOrDefault provides a mock function with given fields: ctx, filters
func (_m *GenericRepository[TEntity]) FirstOrDefault(ctx context.Context, filters map[string]interface{}) (TEntity, error) {
	ret := _m.Called(ctx, filters)

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

// GenericRepository_FirstOrDefault_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FirstOrDefault'
type GenericRepository_FirstOrDefault_Call[TEntity interface{}] struct {
	*mock.Call
}

// FirstOrDefault is a helper method to define mock.On call
//   - ctx context.Context
//   - filters map[string]interface{}
func (_e *GenericRepository_Expecter[TEntity]) FirstOrDefault(ctx interface{}, filters interface{}) *GenericRepository_FirstOrDefault_Call[TEntity] {
	return &GenericRepository_FirstOrDefault_Call[TEntity]{Call: _e.mock.On("FirstOrDefault", ctx, filters)}
}

func (_c *GenericRepository_FirstOrDefault_Call[TEntity]) Run(run func(ctx context.Context, filters map[string]interface{})) *GenericRepository_FirstOrDefault_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]interface{}))
	})
	return _c
}

func (_c *GenericRepository_FirstOrDefault_Call[TEntity]) Return(_a0 TEntity, _a1 error) *GenericRepository_FirstOrDefault_Call[TEntity] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GenericRepository_FirstOrDefault_Call[TEntity]) RunAndReturn(run func(context.Context, map[string]interface{}) (TEntity, error)) *GenericRepository_FirstOrDefault_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: ctx, listQuery
func (_m *GenericRepository[TEntity]) GetAll(ctx context.Context, listQuery *utils.ListQuery) (*utils.ListResult[TEntity], error) {
	ret := _m.Called(ctx, listQuery)

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

// GenericRepository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type GenericRepository_GetAll_Call[TEntity interface{}] struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
//   - listQuery *utils.ListQuery
func (_e *GenericRepository_Expecter[TEntity]) GetAll(ctx interface{}, listQuery interface{}) *GenericRepository_GetAll_Call[TEntity] {
	return &GenericRepository_GetAll_Call[TEntity]{Call: _e.mock.On("GetAll", ctx, listQuery)}
}

func (_c *GenericRepository_GetAll_Call[TEntity]) Run(run func(ctx context.Context, listQuery *utils.ListQuery)) *GenericRepository_GetAll_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*utils.ListQuery))
	})
	return _c
}

func (_c *GenericRepository_GetAll_Call[TEntity]) Return(_a0 *utils.ListResult[TEntity], _a1 error) *GenericRepository_GetAll_Call[TEntity] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GenericRepository_GetAll_Call[TEntity]) RunAndReturn(run func(context.Context, *utils.ListQuery) (*utils.ListResult[TEntity], error)) *GenericRepository_GetAll_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// GetByFilter provides a mock function with given fields: ctx, filters
func (_m *GenericRepository[TEntity]) GetByFilter(ctx context.Context, filters map[string]interface{}) ([]TEntity, error) {
	ret := _m.Called(ctx, filters)

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

// GenericRepository_GetByFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByFilter'
type GenericRepository_GetByFilter_Call[TEntity interface{}] struct {
	*mock.Call
}

// GetByFilter is a helper method to define mock.On call
//   - ctx context.Context
//   - filters map[string]interface{}
func (_e *GenericRepository_Expecter[TEntity]) GetByFilter(ctx interface{}, filters interface{}) *GenericRepository_GetByFilter_Call[TEntity] {
	return &GenericRepository_GetByFilter_Call[TEntity]{Call: _e.mock.On("GetByFilter", ctx, filters)}
}

func (_c *GenericRepository_GetByFilter_Call[TEntity]) Run(run func(ctx context.Context, filters map[string]interface{})) *GenericRepository_GetByFilter_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]interface{}))
	})
	return _c
}

func (_c *GenericRepository_GetByFilter_Call[TEntity]) Return(_a0 []TEntity, _a1 error) *GenericRepository_GetByFilter_Call[TEntity] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GenericRepository_GetByFilter_Call[TEntity]) RunAndReturn(run func(context.Context, map[string]interface{}) ([]TEntity, error)) *GenericRepository_GetByFilter_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// GetByFuncFilter provides a mock function with given fields: ctx, filterFunc
func (_m *GenericRepository[TEntity]) GetByFuncFilter(ctx context.Context, filterFunc func(TEntity) bool) ([]TEntity, error) {
	ret := _m.Called(ctx, filterFunc)

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

// GenericRepository_GetByFuncFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByFuncFilter'
type GenericRepository_GetByFuncFilter_Call[TEntity interface{}] struct {
	*mock.Call
}

// GetByFuncFilter is a helper method to define mock.On call
//   - ctx context.Context
//   - filterFunc func(TEntity) bool
func (_e *GenericRepository_Expecter[TEntity]) GetByFuncFilter(ctx interface{}, filterFunc interface{}) *GenericRepository_GetByFuncFilter_Call[TEntity] {
	return &GenericRepository_GetByFuncFilter_Call[TEntity]{Call: _e.mock.On("GetByFuncFilter", ctx, filterFunc)}
}

func (_c *GenericRepository_GetByFuncFilter_Call[TEntity]) Run(run func(ctx context.Context, filterFunc func(TEntity) bool)) *GenericRepository_GetByFuncFilter_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(TEntity) bool))
	})
	return _c
}

func (_c *GenericRepository_GetByFuncFilter_Call[TEntity]) Return(_a0 []TEntity, _a1 error) *GenericRepository_GetByFuncFilter_Call[TEntity] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GenericRepository_GetByFuncFilter_Call[TEntity]) RunAndReturn(run func(context.Context, func(TEntity) bool) ([]TEntity, error)) *GenericRepository_GetByFuncFilter_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *GenericRepository[TEntity]) GetById(ctx context.Context, id uuid.UUID) (TEntity, error) {
	ret := _m.Called(ctx, id)

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

// GenericRepository_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type GenericRepository_GetById_Call[TEntity interface{}] struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *GenericRepository_Expecter[TEntity]) GetById(ctx interface{}, id interface{}) *GenericRepository_GetById_Call[TEntity] {
	return &GenericRepository_GetById_Call[TEntity]{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *GenericRepository_GetById_Call[TEntity]) Run(run func(ctx context.Context, id uuid.UUID)) *GenericRepository_GetById_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *GenericRepository_GetById_Call[TEntity]) Return(_a0 TEntity, _a1 error) *GenericRepository_GetById_Call[TEntity] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GenericRepository_GetById_Call[TEntity]) RunAndReturn(run func(context.Context, uuid.UUID) (TEntity, error)) *GenericRepository_GetById_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: ctx, searchTerm, listQuery
func (_m *GenericRepository[TEntity]) Search(ctx context.Context, searchTerm string, listQuery *utils.ListQuery) (*utils.ListResult[TEntity], error) {
	ret := _m.Called(ctx, searchTerm, listQuery)

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

// GenericRepository_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type GenericRepository_Search_Call[TEntity interface{}] struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - ctx context.Context
//   - searchTerm string
//   - listQuery *utils.ListQuery
func (_e *GenericRepository_Expecter[TEntity]) Search(ctx interface{}, searchTerm interface{}, listQuery interface{}) *GenericRepository_Search_Call[TEntity] {
	return &GenericRepository_Search_Call[TEntity]{Call: _e.mock.On("Search", ctx, searchTerm, listQuery)}
}

func (_c *GenericRepository_Search_Call[TEntity]) Run(run func(ctx context.Context, searchTerm string, listQuery *utils.ListQuery)) *GenericRepository_Search_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*utils.ListQuery))
	})
	return _c
}

func (_c *GenericRepository_Search_Call[TEntity]) Return(_a0 *utils.ListResult[TEntity], _a1 error) *GenericRepository_Search_Call[TEntity] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GenericRepository_Search_Call[TEntity]) RunAndReturn(run func(context.Context, string, *utils.ListQuery) (*utils.ListResult[TEntity], error)) *GenericRepository_Search_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// SkipTake provides a mock function with given fields: ctx, skip, take
func (_m *GenericRepository[TEntity]) SkipTake(ctx context.Context, skip int, take int) ([]TEntity, error) {
	ret := _m.Called(ctx, skip, take)

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

// GenericRepository_SkipTake_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SkipTake'
type GenericRepository_SkipTake_Call[TEntity interface{}] struct {
	*mock.Call
}

// SkipTake is a helper method to define mock.On call
//   - ctx context.Context
//   - skip int
//   - take int
func (_e *GenericRepository_Expecter[TEntity]) SkipTake(ctx interface{}, skip interface{}, take interface{}) *GenericRepository_SkipTake_Call[TEntity] {
	return &GenericRepository_SkipTake_Call[TEntity]{Call: _e.mock.On("SkipTake", ctx, skip, take)}
}

func (_c *GenericRepository_SkipTake_Call[TEntity]) Run(run func(ctx context.Context, skip int, take int)) *GenericRepository_SkipTake_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *GenericRepository_SkipTake_Call[TEntity]) Return(_a0 []TEntity, _a1 error) *GenericRepository_SkipTake_Call[TEntity] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GenericRepository_SkipTake_Call[TEntity]) RunAndReturn(run func(context.Context, int, int) ([]TEntity, error)) *GenericRepository_SkipTake_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, entity
func (_m *GenericRepository[TEntity]) Update(ctx context.Context, entity TEntity) error {
	ret := _m.Called(ctx, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, TEntity) error); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenericRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type GenericRepository_Update_Call[TEntity interface{}] struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - entity TEntity
func (_e *GenericRepository_Expecter[TEntity]) Update(ctx interface{}, entity interface{}) *GenericRepository_Update_Call[TEntity] {
	return &GenericRepository_Update_Call[TEntity]{Call: _e.mock.On("Update", ctx, entity)}
}

func (_c *GenericRepository_Update_Call[TEntity]) Run(run func(ctx context.Context, entity TEntity)) *GenericRepository_Update_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(TEntity))
	})
	return _c
}

func (_c *GenericRepository_Update_Call[TEntity]) Return(_a0 error) *GenericRepository_Update_Call[TEntity] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GenericRepository_Update_Call[TEntity]) RunAndReturn(run func(context.Context, TEntity) error) *GenericRepository_Update_Call[TEntity] {
	_c.Call.Return(run)
	return _c
}

// UpdateAll provides a mock function with given fields: ctx, entities
func (_m *GenericRepository[TEntity]) UpdateAll(ctx context.Context, entities []TEntity) error {
	ret := _m.Called(ctx, entities)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []TEntity) error); ok {
		r0 = rf(ctx, entities)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenericRepository_UpdateAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAll'
type GenericRepository_UpdateAll_Call[TEntity interface{}] struct {
	*mock.Call
}

// UpdateAll is a helper method to define mock.On call
//   - ctx context.Context
//   - entities []TEntity
func (_e *GenericRepository_Expecter[TEntity]) UpdateAll(ctx interface{}, entities interface{}) *GenericRepository_UpdateAll_Call[TEntity] {
	return &GenericRepository_UpdateAll_Call[TEntity]{Call: _e.mock.On("UpdateAll", ctx, entities)}
}

func (_c *GenericRepository_UpdateAll_Call[TEntity]) Run(run func(ctx context.Context, entities []TEntity)) *GenericRepository_UpdateAll_Call[TEntity] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]TEntity))
	})
	return _c
}

func (_c *GenericRepository_UpdateAll_Call[TEntity]) Return(_a0 error) *GenericRepository_UpdateAll_Call[TEntity] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GenericRepository_UpdateAll_Call[TEntity]) RunAndReturn(run func(context.Context, []TEntity) error) *GenericRepository_UpdateAll_Call[TEntity] {
	_c.Call.Return(run)
	return _c
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
