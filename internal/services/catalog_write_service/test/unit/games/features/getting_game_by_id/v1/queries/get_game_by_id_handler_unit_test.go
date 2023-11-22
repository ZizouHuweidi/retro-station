//go:build unit
// +build unit

package queries

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"

	getGameByIdQueryV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_game_by_id/v1/queries"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type getGameByIdHandlerTest struct {
	*unit_test.UnitTestSharedFixture
}

func TestGetGameByIdHandlerUnit(t *testing.T) {
	suite.Run(
		t,
		&getGameByIdHandlerTest{UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t)},
	)
}

func (c *getGameByIdHandlerTest) Test_Get_Game_By_Id() {
	game := c.Items[0]
	id := uuid.NewV4()
	testCases := []struct {
		Name                       string
		id                         uuid.UUID
		HandlerError               error
		GameRepositoryNumberOfCall int
		ExpectedName               string
		ExpectedId                 uuid.UUID
		RepositoryReturnGame       *models.Game
		RepositoryReturnError      error
		fn                         func()
	}{
		{
			Name:                       "Handle_Should_Get_Game_Successfully",
			id:                         game.GameId,
			HandlerError:               nil,
			GameRepositoryNumberOfCall: 1,
			ExpectedId:                 game.GameId,
			ExpectedName:               game.Name,
			RepositoryReturnGame:       game,
			RepositoryReturnError:      nil,
		},
		{
			Name: "Handle_Should_Return_NotFound_Error_For_NotFound_Item",
			id:   id,
			HandlerError: customErrors.NewApplicationErrorWithCode(
				fmt.Sprintf("error in getting game with id %s in the repository", id.String()),
				http.StatusNotFound,
			),
			GameRepositoryNumberOfCall: 1,
			ExpectedId:                 *new(uuid.UUID),
			ExpectedName:               "",
			RepositoryReturnGame:       nil,
			RepositoryReturnError:      customErrors.NewNotFoundError("game not found"),
		},
		{
			Name: "Handle_Should_Return_Error_For_Error_In_Mapping",
			id:   game.GameId,
			HandlerError: customErrors.NewApplicationErrorWithCode(
				"error in the mapping game",
				http.StatusInternalServerError,
			),
			GameRepositoryNumberOfCall: 1,
			ExpectedId:                 *new(uuid.UUID),
			ExpectedName:               "",
			RepositoryReturnGame:       game,
			RepositoryReturnError:      nil,
			fn: func() {
				mapper.ClearMappings()
			},
		},
	}

	ctx := context.Background()
	for _, testCase := range testCases {
		c.Run(testCase.Name, func() {
			// arrange
			// create new mocks or clear mocks before executing
			c.CleanupMocks()

			getGameByIdHandler := getGameByIdQueryV1.NewGetGameByIdHandler(
				c.Log,
				c.GameRepository,
				c.Tracer,
			)

			c.GameRepository.On("GetGameById", mock.Anything, testCase.id).
				Once().
				Return(testCase.RepositoryReturnGame, testCase.RepositoryReturnError)

			if testCase.fn != nil {
				testCase.fn()
			}

			query, err := getGameByIdQueryV1.NewGetGameById(testCase.id)
			c.Require().NoError(err)

			// act
			dto, err := getGameByIdHandler.Handle(ctx, query)

			// assert
			c.GameRepository.AssertNumberOfCalls(
				c.T(),
				"GetGameById",
				testCase.GameRepositoryNumberOfCall,
			)
			if testCase.HandlerError == nil {
				// success path with a valid dto
				c.Require().NoError(err)
				c.NotNil(dto.Game)
				c.Equal(testCase.ExpectedId, dto.Game.GameId)
				c.Equal(testCase.ExpectedName, dto.Game.Name)
			} else {
				// handler error path
				c.Nil(dto)
				c.ErrorContains(err, testCase.HandlerError.Error())
				if customErrors.IsApplicationError(testCase.HandlerError, http.StatusNotFound) {
					// not found error
					c.True(customErrors.IsNotFoundError(err))
					c.True(customErrors.IsApplicationError(err, http.StatusNotFound))
					c.ErrorContains(err, testCase.HandlerError.Error())
				} else {
					// mapping error
					c.ErrorContains(err, testCase.HandlerError.Error())
					c.True(customErrors.IsApplicationError(err, http.StatusInternalServerError))
				}
			}
		})
	}

	//c.Register("Handle_Should_Get_Game_Successfully", func() {
	//	// create new mocks or clear mocks before executing
	//	c.UnitTestMockFixture = unit_test.NewUnitTestMockFixture(c.T())
	//	c.getGameByIdHandler = NewGetGameByIdHandler(c.Log, c.Cfg, c.GameRepository)
	//
	//	c.GameRepository.On("GetGameById", mock.Anything, game.GameId).
	//		Once().
	//		Return(game, nil)
	//
	//	query := NewGetGameById(game.GameId)
	//
	//	dto, err := c.getGameByIdHandler.Handle(c.Ctx, query)
	//	c.Require().NoError(err)
	//
	//	c.GameRepository.AssertNumberOfCalls(c.T(), "GetGameById", 1)
	//	c.Equal(game.GameId, dto.Game.GameId)
	//	c.Equal(game.Name, dto.Game.Name)
	//})
	//
	//c.Register("Handle_Should_Return_NotFound_Error_For_NotFound_Item", func() {
	//	// create new mocks or clear mocks before executing
	//	c.UnitTestMockFixture = unit_test.NewUnitTestMockFixture(c.T())
	//	c.getGameByIdHandler = NewGetGameByIdHandler(c.Log, c.Cfg, c.GameRepository)
	//
	//	c.GameRepository.On("GetGameById", mock.Anything, id).
	//		Once().
	//		Return(nil, customErrors.NewNotFoundError("game not found"))
	//
	//	query := NewGetGameById(id)
	//
	//	dto, err := c.getGameByIdHandler.Handle(c.Ctx, query)
	//
	//	c.Require().Error(err)
	//	c.True(customErrors.IsApplicationError(err, http.StatusNotFound))
	//	c.True(customErrors.IsNotFoundError(err))
	//	c.ErrorContains(err, fmt.Sprintf("error in getting game with id %s in the repository", id.String()))
	//	c.Nil(dto)
	//
	//	c.GameRepository.AssertNumberOfCalls(c.T(), "GetGameById", 1)
	//})
	//
	//c.Register("Handle_Should_Return_Error_For_Error_In_Mapping", func() {
	//	// create new mocks or clear mocks before executing
	//	c.UnitTestMockFixture = unit_test.NewUnitTestMockFixture(c.T())
	//	c.getGameByIdHandler = NewGetGameByIdHandler(c.Log, c.Cfg, c.GameRepository)
	//
	//	game := testData.Games[0]
	//	c.GameRepository.On("GetGameById", mock.Anything, game.GameId).
	//		Once().
	//		Return(game, nil)
	//
	//	mapper.ClearMappings()
	//
	//	query := NewGetGameById(game.GameId)
	//
	//	dto, err := c.getGameByIdHandler.Handle(c.Ctx, query)
	//
	//	c.GameRepository.AssertNumberOfCalls(c.T(), "GetGameById", 1)
	//	c.Nil(dto)
	//	c.Require().Error(err)
	//	c.True(customErrors.IsApplicationError(err, http.StatusInternalServerError))
	//	c.ErrorContains(err, "error in the mapping game")
	//})
}
