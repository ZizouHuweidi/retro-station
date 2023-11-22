//go:build unit
// +build unit

package commands

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type updateGameUnitTests struct {
	*unit_test.UnitTestSharedFixture
}

func TestUpdateGameUnit(t *testing.T) {
	suite.Run(
		t,
		&updateGameUnitTests{UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t)},
	)
}

func (c *updateGameUnitTests) Test_New_Update_Game_Should_Return_No_Error_For_Valid_Input() {
	id := uuid.NewV4()
	name := gofakeit.Name()
	description := gofakeit.EmojiDescription()
	price := gofakeit.Price(150, 6000)

	updateGame, err := commands.NewUpdateGame(id, name, description, price)

	c.Assert().NotNil(updateGame)
	c.Assert().Equal(id, updateGame.GameID)
	c.Assert().Equal(name, updateGame.Name)
	c.Assert().Equal(price, updateGame.Price)

	c.Require().NoError(err)
}

func (c *updateGameUnitTests) Test_New_Update_Game_Should_Return_Error_For_Invalid_Price() {
	command, err := commands.NewUpdateGame(
		uuid.NewV4(),
		gofakeit.Name(),
		gofakeit.EmojiDescription(),
		0,
	)

	c.Require().Error(err)
	c.Assert().Nil(command)
}

func (c *updateGameUnitTests) Test_New_Update_Game_Should_Return_Error_For_Empty_Name() {
	command, err := commands.NewUpdateGame(uuid.NewV4(), "", gofakeit.EmojiDescription(), 120)

	c.Require().Error(err)
	c.Assert().Nil(command)
}

func (c *updateGameUnitTests) Test_New_Update_Game_Should_Return_Error_For_Empty_Description() {
	command, err := commands.NewUpdateGame(uuid.NewV4(), gofakeit.Name(), "", 120)

	c.Require().Error(err)
	c.Assert().Nil(command)
}
