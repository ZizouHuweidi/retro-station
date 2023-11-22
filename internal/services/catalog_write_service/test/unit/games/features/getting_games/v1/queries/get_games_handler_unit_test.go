//go:build unit
// +build unit

package queries

import (
	"context"
	"net/http"
	"testing"

	"emperror.dev/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_games/v1/queries"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/mocks/testData"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type getGamesHandlerUnitTests struct {
	*unit_test.UnitTestSharedFixture
}

func TestGetGamesUnit(t *testing.T) {
	suite.Run(
		t,
		&getGamesHandlerUnitTests{UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t)},
	)
}

func (c *getGamesHandlerUnitTests) Test_Handle_Should_Return_Games_Successfully() {
	ctx := context.Background()

	query, err := queries.NewGetGames(utils.NewListQuery(10, 1))
	c.Require().NoError(err)

	items := utils.NewListResult[*models.Game](
		testData.Games,
		10,
		1,
		int64(len(testData.Games)),
	)
	c.GameRepository.On("GetAllGames", mock.Anything, mock.Anything).
		Once().
		Return(items, nil)

	getGamesHandler := queries.NewGetGamesHandler(c.Log, c.GameRepository, c.Tracer)

	res, err := getGamesHandler.Handle(ctx, query)
	c.Require().NoError(err)
	c.NotNil(res)
	c.NotEmpty(res.Games)
	c.Equal(len(testData.Games), len(res.Games.Items))
	c.GameRepository.AssertNumberOfCalls(c.T(), "GetAllGames", 1)
}

func (c *getGamesHandlerUnitTests) Test_Handle_Should_Return_Error_For_Error_In_Fetching_Data_From_Repository() {
	ctx := context.Background()

	query, err := queries.NewGetGames(utils.NewListQuery(10, 1))
	c.Require().NoError(err)

	c.GameRepository.On("GetAllGames", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("error in fetching games from repository"))

	getGamesHandler := queries.NewGetGamesHandler(c.Log, c.GameRepository, c.Tracer)

	res, err := getGamesHandler.Handle(ctx, query)
	c.Require().Error(err)
	c.True(customErrors.IsApplicationError(err, http.StatusInternalServerError))
	c.Nil(res)
	c.GameRepository.AssertNumberOfCalls(c.T(), "GetAllGames", 1)
}

func (c *getGamesHandlerUnitTests) Test_Handle_Should_Return_Error_For_Mapping_List_Result() {
	ctx := context.Background()

	query, err := queries.NewGetGames(utils.NewListQuery(10, 1))
	c.Require().NoError(err)

	c.GameRepository.On("GetAllGames", mock.Anything, mock.Anything).
		Once().
		Return(nil, nil)

	getGamesHandler := queries.NewGetGamesHandler(c.Log, c.GameRepository, c.Tracer)

	res, err := getGamesHandler.Handle(ctx, query)
	c.Require().Error(err)
	c.True(customErrors.IsApplicationError(err, http.StatusInternalServerError))
	c.Nil(res)
	c.GameRepository.AssertNumberOfCalls(c.T(), "GetAllGames", 1)
}
