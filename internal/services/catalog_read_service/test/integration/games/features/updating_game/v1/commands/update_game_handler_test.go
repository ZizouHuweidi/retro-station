//go:build integration
// +build integration

package commands

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mehdihadeli/go-mediatr"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/updating_games/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestUpdateGame(t *testing.T) {
	integrationTestSharedFixture := integration.NewIntegrationTestSharedFixture(t)

	Convey("Updating Game Feature", t, func() {
		ctx := context.Background()
		integrationTestSharedFixture.InitializeTest()

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey("Updating an existing game in the database", func() {
			Convey("Given an existing game in the system", func() {
				gameId, err := uuid.FromString(integrationTestSharedFixture.Items[0].GameId)
				So(err, ShouldBeNil)

				updateGame, err := commands.NewUpdateGame(
					gameId,
					gofakeit.Name(),
					gofakeit.AdjectiveDescriptive(),
					gofakeit.Price(150, 6000),
				)
				So(err, ShouldBeNil)

				Convey("When a UpdateGame command executed for a existing game", func() {
					result, err := mediatr.Send[*commands.UpdateGame, *mediatr.Unit](ctx, updateGame)

					Convey("Then the game should be updated successfully", func() {
						// Assert that the error is nil (indicating success).
						So(err, ShouldBeNil)
						So(result, ShouldNotBeNil)

						Convey("And the updated game details should be reflected in the system", func() {
							// Fetch the updated game from the database.
							updatedGame, _ := integrationTestSharedFixture.GameRepository.GetGameByGameId(
								ctx,
								gameId.String(),
							)

							Convey("And the game's properties should match the updated data", func() {
								// Assert that the game properties match the updated data.
								So(updatedGame.Name, ShouldEqual, updatedGame.Name)
								So(updatedGame.Price, ShouldEqual, updatedGame.Price)
								// Add more assertions as needed for other properties.
							})
						})
					})
				})
			})
		})

		integrationTestSharedFixture.DisposeTest()
	})
}
