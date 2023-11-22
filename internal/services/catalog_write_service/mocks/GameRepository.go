// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"

	utils "github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	uuid "github.com/satori/go.uuid"
)

// GameRepository is an autogenerated mock type for the GameRepository type
type GameRepository struct {
	mock.Mock
}

type GameRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *GameRepository) EXPECT() *GameRepository_Expecter {
	return &GameRepository_Expecter{mock: &_m.Mock}
}

// CreateGame provides a mock function with given fields: ctx, game
func (_m *GameRepository) CreateGame(ctx context.Context, game *models.Game) (*models.Game, error) {
	ret := _m.Called(ctx, game)

	var r0 *models.Game
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Game) (*models.Game, error)); ok {
		return rf(ctx, game)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Game) *models.Game); ok {
		r0 = rf(ctx, game)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Game)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Game) error); ok {
		r1 = rf(ctx, game)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameRepository_CreateGame_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateGame'
type GameRepository_CreateGame_Call struct {
	*mock.Call
}

// CreateGame is a helper method to define mock.On call
//   - ctx context.Context
//   - game *models.Game
func (_e *GameRepository_Expecter) CreateGame(ctx interface{}, game interface{}) *GameRepository_CreateGame_Call {
	return &GameRepository_CreateGame_Call{Call: _e.mock.On("CreateGame", ctx, game)}
}

func (_c *GameRepository_CreateGame_Call) Run(run func(ctx context.Context, game *models.Game)) *GameRepository_CreateGame_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Game))
	})
	return _c
}

func (_c *GameRepository_CreateGame_Call) Return(_a0 *models.Game, _a1 error) *GameRepository_CreateGame_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GameRepository_CreateGame_Call) RunAndReturn(run func(context.Context, *models.Game) (*models.Game, error)) *GameRepository_CreateGame_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteGameByID provides a mock function with given fields: ctx, _a1
func (_m *GameRepository) DeleteGameByID(ctx context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GameRepository_DeleteGameByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteGameByID'
type GameRepository_DeleteGameByID_Call struct {
	*mock.Call
}

// DeleteGameByID is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 uuid.UUID
func (_e *GameRepository_Expecter) DeleteGameByID(ctx interface{}, _a1 interface{}) *GameRepository_DeleteGameByID_Call {
	return &GameRepository_DeleteGameByID_Call{Call: _e.mock.On("DeleteGameByID", ctx, _a1)}
}

func (_c *GameRepository_DeleteGameByID_Call) Run(run func(ctx context.Context, _a1 uuid.UUID)) *GameRepository_DeleteGameByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *GameRepository_DeleteGameByID_Call) Return(_a0 error) *GameRepository_DeleteGameByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GameRepository_DeleteGameByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *GameRepository_DeleteGameByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllGames provides a mock function with given fields: ctx, listQuery
func (_m *GameRepository) GetAllGames(ctx context.Context, listQuery *utils.ListQuery) (*utils.ListResult[*models.Game], error) {
	ret := _m.Called(ctx, listQuery)

	var r0 *utils.ListResult[*models.Game]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *utils.ListQuery) (*utils.ListResult[*models.Game], error)); ok {
		return rf(ctx, listQuery)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *utils.ListQuery) *utils.ListResult[*models.Game]); ok {
		r0 = rf(ctx, listQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ListResult[*models.Game])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *utils.ListQuery) error); ok {
		r1 = rf(ctx, listQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameRepository_GetAllGames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllGames'
type GameRepository_GetAllGames_Call struct {
	*mock.Call
}

// GetAllGames is a helper method to define mock.On call
//   - ctx context.Context
//   - listQuery *utils.ListQuery
func (_e *GameRepository_Expecter) GetAllGames(ctx interface{}, listQuery interface{}) *GameRepository_GetAllGames_Call {
	return &GameRepository_GetAllGames_Call{Call: _e.mock.On("GetAllGames", ctx, listQuery)}
}

func (_c *GameRepository_GetAllGames_Call) Run(run func(ctx context.Context, listQuery *utils.ListQuery)) *GameRepository_GetAllGames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*utils.ListQuery))
	})
	return _c
}

func (_c *GameRepository_GetAllGames_Call) Return(_a0 *utils.ListResult[*models.Game], _a1 error) *GameRepository_GetAllGames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GameRepository_GetAllGames_Call) RunAndReturn(run func(context.Context, *utils.ListQuery) (*utils.ListResult[*models.Game], error)) *GameRepository_GetAllGames_Call {
	_c.Call.Return(run)
	return _c
}

// GetGameById provides a mock function with given fields: ctx, _a1
func (_m *GameRepository) GetGameById(ctx context.Context, _a1 uuid.UUID) (*models.Game, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *models.Game
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*models.Game, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *models.Game); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Game)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameRepository_GetGameById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGameById'
type GameRepository_GetGameById_Call struct {
	*mock.Call
}

// GetGameById is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 uuid.UUID
func (_e *GameRepository_Expecter) GetGameById(ctx interface{}, _a1 interface{}) *GameRepository_GetGameById_Call {
	return &GameRepository_GetGameById_Call{Call: _e.mock.On("GetGameById", ctx, _a1)}
}

func (_c *GameRepository_GetGameById_Call) Run(run func(ctx context.Context, _a1 uuid.UUID)) *GameRepository_GetGameById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *GameRepository_GetGameById_Call) Return(_a0 *models.Game, _a1 error) *GameRepository_GetGameById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GameRepository_GetGameById_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*models.Game, error)) *GameRepository_GetGameById_Call {
	_c.Call.Return(run)
	return _c
}

// SearchGames provides a mock function with given fields: ctx, searchText, listQuery
func (_m *GameRepository) SearchGames(ctx context.Context, searchText string, listQuery *utils.ListQuery) (*utils.ListResult[*models.Game], error) {
	ret := _m.Called(ctx, searchText, listQuery)

	var r0 *utils.ListResult[*models.Game]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *utils.ListQuery) (*utils.ListResult[*models.Game], error)); ok {
		return rf(ctx, searchText, listQuery)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *utils.ListQuery) *utils.ListResult[*models.Game]); ok {
		r0 = rf(ctx, searchText, listQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ListResult[*models.Game])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *utils.ListQuery) error); ok {
		r1 = rf(ctx, searchText, listQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameRepository_SearchGames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchGames'
type GameRepository_SearchGames_Call struct {
	*mock.Call
}

// SearchGames is a helper method to define mock.On call
//   - ctx context.Context
//   - searchText string
//   - listQuery *utils.ListQuery
func (_e *GameRepository_Expecter) SearchGames(ctx interface{}, searchText interface{}, listQuery interface{}) *GameRepository_SearchGames_Call {
	return &GameRepository_SearchGames_Call{Call: _e.mock.On("SearchGames", ctx, searchText, listQuery)}
}

func (_c *GameRepository_SearchGames_Call) Run(run func(ctx context.Context, searchText string, listQuery *utils.ListQuery)) *GameRepository_SearchGames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*utils.ListQuery))
	})
	return _c
}

func (_c *GameRepository_SearchGames_Call) Return(_a0 *utils.ListResult[*models.Game], _a1 error) *GameRepository_SearchGames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GameRepository_SearchGames_Call) RunAndReturn(run func(context.Context, string, *utils.ListQuery) (*utils.ListResult[*models.Game], error)) *GameRepository_SearchGames_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateGame provides a mock function with given fields: ctx, game
func (_m *GameRepository) UpdateGame(ctx context.Context, game *models.Game) (*models.Game, error) {
	ret := _m.Called(ctx, game)

	var r0 *models.Game
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Game) (*models.Game, error)); ok {
		return rf(ctx, game)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Game) *models.Game); ok {
		r0 = rf(ctx, game)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Game)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Game) error); ok {
		r1 = rf(ctx, game)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameRepository_UpdateGame_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateGame'
type GameRepository_UpdateGame_Call struct {
	*mock.Call
}

// UpdateGame is a helper method to define mock.On call
//   - ctx context.Context
//   - game *models.Game
func (_e *GameRepository_Expecter) UpdateGame(ctx interface{}, game interface{}) *GameRepository_UpdateGame_Call {
	return &GameRepository_UpdateGame_Call{Call: _e.mock.On("UpdateGame", ctx, game)}
}

func (_c *GameRepository_UpdateGame_Call) Run(run func(ctx context.Context, game *models.Game)) *GameRepository_UpdateGame_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Game))
	})
	return _c
}

func (_c *GameRepository_UpdateGame_Call) Return(_a0 *models.Game, _a1 error) *GameRepository_UpdateGame_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GameRepository_UpdateGame_Call) RunAndReturn(run func(context.Context, *models.Game) (*models.Game, error)) *GameRepository_UpdateGame_Call {
	_c.Call.Return(run)
	return _c
}

// NewGameRepository creates a new instance of GameRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGameRepository(t interface {
	mock.TestingT
	Cleanup(func())
},
) *GameRepository {
	mock := &GameRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}