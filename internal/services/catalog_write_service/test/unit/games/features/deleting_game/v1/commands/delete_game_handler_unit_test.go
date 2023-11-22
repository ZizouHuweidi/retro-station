//go:build unit
// +build unit

package commands

import (
	"context"
	"net/http"
	"testing"

	"emperror.dev/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/deleting_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type deleteGameHandlerUnitTests struct {
	*unit_test.UnitTestSharedFixture
}

func TestDeleteGameHandlerUnit(t *testing.T) {
	suite.Run(
		t,
		&deleteGameHandlerUnitTests{
			UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t),
		},
	)
}

func (c *deleteGameHandlerUnitTests) Test_Handle_Should_Delete_Game_With_Valid_Id() {
	ctx := context.Background()
	id := c.Items[0].GameId

	deleteGame := &commands.DeleteGame{
		GameID: id,
	}

	c.GameRepository.On("DeleteGameByID", mock.Anything, id).
		Once().
		Return(nil)

	deleteGameHandler := commands.NewDeleteGameHandler(c.Log, c.Uow, c.Bus, c.Tracer)

	_, err := deleteGameHandler.Handle(ctx, deleteGame)
	c.Require().NoError(err)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "DeleteGameByID", 1)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 1)
}

func (c *deleteGameHandlerUnitTests) Test_Handle_Should_Return_NotFound_Error_When_Id_Is_Invalid() {
	ctx := context.Background()
	id := uuid.NewV4()

	deleteGame := &commands.DeleteGame{
		GameID: id,
	}

	c.GameRepository.On("DeleteGameByID", mock.Anything, id).
		Once().
		Return(customErrors.NewNotFoundError("error finding game"))

	deleteGameHandler := commands.NewDeleteGameHandler(c.Log, c.Uow, c.Bus, c.Tracer)

	res, err := deleteGameHandler.Handle(ctx, deleteGame)
	c.Require().Error(err)
	c.True(customErrors.IsNotFoundError(err))
	c.True(customErrors.IsApplicationError(err, http.StatusNotFound))
	c.NotNil(res)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "DeleteGameByID", 1)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 0)
}

func (c *deleteGameHandlerUnitTests) Test_Handle_Should_Return_Error_For_Error_In_Bus() {
	ctx := context.Background()
	id := c.Items[0].GameId

	deleteGame := &commands.DeleteGame{
		GameID: id,
	}

	c.GameRepository.On("DeleteGameByID", mock.Anything, id).
		Once().
		Return(nil)

	// override called mock
	// https://github.com/stretchr/testify/issues/558
	c.Bus.Mock.ExpectedCalls = nil
	c.Bus.On("PublishMessage", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(errors.New("error in the publish message"))

	deleteGameHandler := commands.NewDeleteGameHandler(c.Log, c.Uow, c.Bus, c.Tracer)
	dto, err := deleteGameHandler.Handle(ctx, deleteGame)

	c.NotNil(dto)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "DeleteGameByID", 1)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 1)
	c.True(customErrors.IsApplicationError(err, http.StatusInternalServerError))
	c.ErrorContains(err, "error in publishing 'GameDeleted' message")
}
