//go:build unit
// +build unit

package commands

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"

	createGameCommandV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/mocks/testData"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type createGameHandlerUnitTests struct {
	*unit_test.UnitTestSharedFixture
}

func TestCreateGameHandlerUnit(t *testing.T) {
	suite.Run(
		t,
		&createGameHandlerUnitTests{
			UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t),
		},
	)
}

func (c *createGameHandlerUnitTests) Test_Handle_Should_Create_New_Game_With_Valid_Data() {
	ctx := context.Background()
	id := uuid.NewV4()

	createGameHandler := createGameCommandV1.NewCreateGameHandler(
		c.Log,
		c.Uow,
		c.Bus,
		c.Tracer,
	)

	createGame := &createGameCommandV1.CreateGame{
		GameID:      id,
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.EmojiDescription(),
		Price:       gofakeit.Price(100, 1000),
	}

	game := testData.Games[0]

	c.GameRepository.On("CreateGame", mock.Anything, mock.Anything).
		Once().
		Return(game, nil)

	dto, err := createGameHandler.Handle(ctx, createGame)
	c.Require().NoError(err)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "CreateGame", 1)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 1)
	c.Equal(dto.GameID, id)
}

func (c *createGameHandlerUnitTests) Test_Handle_Should_Return_Error_For_Duplicate_Item() {
	ctx := context.Background()
	id := uuid.NewV4()

	createGame := &createGameCommandV1.CreateGame{
		GameID:      id,
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.EmojiDescription(),
		Price:       gofakeit.Price(100, 1000),
	}

	c.GameRepository.On("CreateGame", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("error duplicate game"))

	createGameHandler := createGameCommandV1.NewCreateGameHandler(
		c.Log,
		c.Uow,
		c.Bus,
		c.Tracer,
	)

	dto, err := createGameHandler.Handle(ctx, createGame)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "CreateGame", 1)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 0)
	c.True(customErrors.IsApplicationError(err, http.StatusConflict))
	c.ErrorContains(err, "game already exists")
	c.Nil(dto)
}

func (c *createGameHandlerUnitTests) Test_Handle_Should_Return_Error_For_Error_In_Bus() {
	ctx := context.Background()
	id := uuid.NewV4()

	createGame := &createGameCommandV1.CreateGame{
		GameID:      id,
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.EmojiDescription(),
		Price:       gofakeit.Price(100, 1000),
	}

	game := testData.Games[0]
	c.GameRepository.On("CreateGame", mock.Anything, mock.Anything).
		Once().
		Return(game, nil)

	// override called mock
	// https://github.com/stretchr/testify/issues/558
	c.Bus.Mock.ExpectedCalls = nil
	c.Bus.On("PublishMessage", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(errors.New("error in the publish message"))

	createGameHandler := createGameCommandV1.NewCreateGameHandler(
		c.Log,
		c.Uow,
		c.Bus,
		c.Tracer,
	)

	dto, err := createGameHandler.Handle(ctx, createGame)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "CreateGame", 1)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 1)
	c.ErrorContains(err, "error in the publish message")
	c.ErrorContains(err, "error in publishing GameCreated integration_events event")
	c.Nil(dto)
}

func (c *createGameHandlerUnitTests) Test_Handle_Should_Return_Error_For_Error_In_Mapping() {
	ctx := context.Background()
	id := uuid.NewV4()

	createGame := &createGameCommandV1.CreateGame{
		GameID:      id,
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.EmojiDescription(),
		Price:       gofakeit.Price(100, 1000),
	}

	game := testData.Games[0]
	c.GameRepository.On("CreateGame", mock.Anything, mock.Anything).
		Once().
		Return(game, nil)

	mapper.ClearMappings()

	createGameHandler := createGameCommandV1.NewCreateGameHandler(
		c.Log,
		c.Uow,
		c.Bus,
		c.Tracer,
	)

	dto, err := createGameHandler.Handle(ctx, createGame)

	c.Uow.AssertNumberOfCalls(c.T(), "Do", 1)
	c.GameRepository.AssertNumberOfCalls(c.T(), "CreateGame", 1)
	c.Bus.AssertNumberOfCalls(c.T(), "PublishMessage", 0)
	c.ErrorContains(err, "error in the mapping GameDto")
	c.True(customErrors.IsApplicationError(err, http.StatusInternalServerError))
	c.Nil(dto)
}
