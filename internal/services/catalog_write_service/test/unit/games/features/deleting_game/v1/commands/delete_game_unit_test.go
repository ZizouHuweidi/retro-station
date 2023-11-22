//go:build unit
// +build unit

package commands

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/deleting_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type deleteGameUnitTests struct {
	*unit_test.UnitTestSharedFixture
}

func TestDeleteGameByIdUnit(t *testing.T) {
	suite.Run(
		t,
		&deleteGameUnitTests{UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t)},
	)
}

func (c *deleteGameUnitTests) Test_New_Delete_Game_Should_Return_No_Error_For_Valid_Input() {
	id := uuid.NewV4()

	query, err := commands.NewDeleteGame(id)

	c.Assert().NotNil(query)
	c.Assert().Equal(query.GameID, id)
	c.Require().NoError(err)
}

func (c *deleteGameUnitTests) Test_New_Delete_Game_Should_Return_Error_For_Invalid_Id() {
	query, err := commands.NewDeleteGame(uuid.UUID{})

	c.Assert().Nil(query)
	c.Require().Error(err)
}
