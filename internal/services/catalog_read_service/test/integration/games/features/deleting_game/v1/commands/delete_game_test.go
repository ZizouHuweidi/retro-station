//go:build integration
// +build integration

package commands

import (
	"context"
	"testing"

	"github.com/mehdihadeli/go-mediatr"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/deleting_games/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestDeleteGame(t *testing.T) {
	integrationTestSharedFixture := integration.NewIntegrationTestSharedFixture(
		t,
	)

	Convey("Deleting Game Feature", t, func() {
		ctx := context.Background()
		integrationTestSharedFixture.InitializeTest()

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey("Deleting an existing game from the database", func() {
			Convey("Given an existing game in the mongo database", func() {
				gameId, err := uuid.FromString(
					integrationTestSharedFixture.Items[0].GameId,
				)
				So(err, ShouldBeNil)

				command, err := commands.NewDeleteGame(gameId)
				So(err, ShouldBeNil)

				Convey("When we execute the DeleteGame command", func() {
					result, err := mediatr.Send[*commands.DeleteGame, *mediatr.Unit](
						context.Background(),
						command,
					)

					Convey(
						"Then the game should be deleted successfully in mongo database",
						func() {
							So(err, ShouldBeNil)
							So(result, ShouldNotBeNil)

							Convey(
								"And the game should no longer exist in the system",
								func() {
									deletedGame, _ := integrationTestSharedFixture.GameRepository.GetGameByGameId(
										ctx,
										gameId.String(),
									)
									So(deletedGame, ShouldBeNil)
								},
							)
						},
					)
				})
			})
		})

		integrationTestSharedFixture.DisposeTest()
	})
}
