// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	data "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
)

// CatalogUnitOfWorkActionFunc is an autogenerated mock type for the CatalogUnitOfWorkActionFunc type
type CatalogUnitOfWorkActionFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: catalogContext
func (_m *CatalogUnitOfWorkActionFunc) Execute(catalogContext data.CatalogContext) error {
	ret := _m.Called(catalogContext)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(data.CatalogContext) error); ok {
		r0 = rf(catalogContext)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCatalogUnitOfWorkActionFunc creates a new instance of CatalogUnitOfWorkActionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCatalogUnitOfWorkActionFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *CatalogUnitOfWorkActionFunc {
	mock := &CatalogUnitOfWorkActionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
