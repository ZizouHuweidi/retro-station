//go:build unit
// +build unit

package queries

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"

	getGameByIdQuery "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_game_by_id/v1/queries"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type getGameByIdUnitTests struct {
	*unit_test.UnitTestSharedFixture
}

func TestGetGameByIdUnit(t *testing.T) {
	suite.Run(
		t,
		&getGameByIdUnitTests{UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t)},
	)
}

func (c *getGameByIdUnitTests) Test_New_Get_Game_By_Id_Should_Return_No_Error_For_Valid_Input() {
	id := uuid.NewV4()

	query, err := getGameByIdQuery.NewGetGameById(id)

	c.Assert().NotNil(query)
	c.Assert().Equal(query.GameID, id)
	c.Require().NoError(err)
}

func (c *getGameByIdUnitTests) Test_New_Get_Game_By_Id_Should_Return_Error_For_Invalid_Id() {
	query, err := getGameByIdQuery.NewGetGameById(uuid.UUID{})

	c.Assert().Nil(query)
	c.Require().Error(err)
}
