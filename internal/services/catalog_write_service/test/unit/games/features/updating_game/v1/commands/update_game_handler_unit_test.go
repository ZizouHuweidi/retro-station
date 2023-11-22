//go:build unit
// +build unit

package commands

import (
	"context"
	"net/http"
	"testing"

	"emperror.dev/errors"
	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type updateGameHandlerUnitTests struct {
	*unit_test.UnitTestSharedFixture
}

func TestUpdateGameHandlerUnit(t *testing.T) {
	suite.Run(
		t,
		&updateGameHandlerUnitTests{
			UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t),
		},
	)
}

func (c *updateGameHandlerUnitTests) Test_Handle_Should_Update_Game_With_Valid_Data() {
	ctx := context.Background()
	existing := c.Items[0]

	updateGameCommand, err := commands.NewUpdateGame(
		existing.GameId,
		gofakeit.Name(),
		gofakeit.EmojiDescription(),
		existing.Price,
	)
	c.Require().NoError(err)

	updated := &models.Game{
		GameId:      existing.GameId,
		Name:        updateGameCommand.Name,
		Description: updateGameCommand.Description,
		UpdatedAt:   updateGameCommand.UpdatedAt,
		CreatedAt:   existing.CreatedAt,
		Price:       existing.Price,
	}

	c.GameRepository.On("GetGameById", mock.Anything, existing.GameId).
		Once().
		Return(existing, nil)

	c.GameRepository.On("UpdateGame", mock.Anything, mock.Anything).
		Once().
		Return(updated, nil)

	updateGameHandler := commands.NewUpdateGameHandler(c.Log, c.Uow, c.Bus, c.Tracer)

	_, err = updateGameHandler.Handle(ctx, updateGameCommand)
	c.Require().NoError(err)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "GetGameById", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "UpdateGame", 1)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 1)
}

func (c *updateGameHandlerUnitTests) Test_Handle_Should_Return_Error_For_NotFound_Item() {
	ctx := context.Background()
	id := uuid.NewV4()

	command, err := commands.NewUpdateGame(
		id,
		gofakeit.Name(),
		gofakeit.EmojiDescription(),
		gofakeit.Price(150, 6000),
	)
	c.Require().NoError(err)

	c.GameRepository.On("GetGameById", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("error notfound game"))

	updateGameHandler := commands.NewUpdateGameHandler(c.Log, c.Uow, c.Bus, c.Tracer)
	dto, err := updateGameHandler.Handle(ctx, command)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "GetGameById", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "UpdateGame", 0)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 0)
	c.True(customErrors.IsApplicationError(err, http.StatusNotFound))
	c.ErrorContains(err, "error notfound game")
	c.NotNil(dto)
}

func (c *updateGameHandlerUnitTests) Test_Handle_Should_Return_Error_For_Error_In_Bus() {
	ctx := context.Background()
	existing := c.Items[0]

	updateGameCommand, err := commands.NewUpdateGame(
		existing.GameId,
		gofakeit.Name(),
		gofakeit.EmojiDescription(),
		existing.Price,
	)
	c.Require().NoError(err)

	updated := &models.Game{
		GameId:      existing.GameId,
		Name:        updateGameCommand.Name,
		Description: updateGameCommand.Description,
		UpdatedAt:   updateGameCommand.UpdatedAt,
		CreatedAt:   existing.CreatedAt,
		Price:       existing.Price,
	}

	c.GameRepository.On("GetGameById", mock.Anything, existing.GameId).
		Once().
		Return(existing, nil)

	c.GameRepository.On("UpdateGame", mock.Anything, mock.Anything).
		Once().
		Return(updated, nil)

	// override called mock
	// https://github.com/stretchr/testify/issues/558
	c.Bus.Mock.ExpectedCalls = nil
	c.Bus.On("PublishMessage", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(errors.New("error in the publish message"))

	updateGameHandler := commands.NewUpdateGameHandler(c.Log, c.Uow, c.Bus, c.Tracer)
	dto, err := updateGameHandler.Handle(ctx, updateGameCommand)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "UpdateGame", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "GetGameById", 1)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 1)
	c.ErrorContains(err, "error in the publish message")
	c.ErrorContains(err, "error in publishing 'GameUpdated' message")
	c.NotNil(dto)
}
