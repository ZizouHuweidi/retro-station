//go:build unit
// +build unit

package commands

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"

	createGameCommand "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type createGameUnitTests struct {
	*unit_test.UnitTestSharedFixture
}

func TestCreateGameUnit(t *testing.T) {
	suite.Run(
		t,
		&createGameUnitTests{UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t)},
	)
}

func (c *createGameUnitTests) Test_New_Create_Game_Should_Return_No_Error_For_Valid_Input() {
	name := gofakeit.Name()
	description := gofakeit.EmojiDescription()
	price := gofakeit.Price(150, 6000)

	updateGame, err := createGameCommand.NewCreateGame(name, description, price)

	c.Assert().NotNil(updateGame)
	c.Assert().Equal(name, updateGame.Name)
	c.Assert().Equal(price, updateGame.Price)

	c.Require().NoError(err)
}

func (c *createGameUnitTests) Test_New_Create_Game_Should_Return_Error_For_Invalid_Price() {
	command, err := createGameCommand.NewCreateGame(
		gofakeit.Name(),
		gofakeit.EmojiDescription(),
		0,
	)

	c.Require().Error(err)
	c.Assert().Nil(command)
}

func (c *createGameUnitTests) Test_New_Create_Game_Should_Return_Error_For_Empty_Name() {
	command, err := createGameCommand.NewCreateGame("", gofakeit.EmojiDescription(), 120)

	c.Require().Error(err)
	c.Assert().Nil(command)
}

func (c *createGameUnitTests) Test_New_Create_Game_Should_Return_Error_For_Empty_Description() {
	command, err := createGameCommand.NewCreateGame(gofakeit.Name(), "", 120)

	c.Require().Error(err)
	c.Assert().Nil(command)
}
