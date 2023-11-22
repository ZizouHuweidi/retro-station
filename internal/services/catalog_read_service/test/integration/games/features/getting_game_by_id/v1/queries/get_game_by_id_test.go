//go:build integration
// +build integration

package queries

import (
	"context"
	"testing"

	"github.com/mehdihadeli/go-mediatr"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/get_game_by_id/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/get_game_by_id/v1/queries"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestGetGameById(t *testing.T) {
	integrationTestSharedFixture := integration.NewIntegrationTestSharedFixture(t)
	ctx := context.Background()

	Convey("Get Game by ID Feature", t, func() {
		integrationTestSharedFixture.InitializeTest()

		knownGameID, err := uuid.FromString(integrationTestSharedFixture.Items[0].Id)
		unknownGameID := uuid.NewV4()
		So(err, ShouldBeNil)

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey("Returning an existing game with valid Id from the database with correct properties", func() {
			Convey("Given a game with a known ID exists in the database", func() {
				query, err := queries.NewGetGameById(knownGameID)
				So(err, ShouldBeNil)

				Convey("When we execute GetGameById query for a game with known ID", func() {
					result, err := mediatr.Send[*queries.GetGameById, *dtos.GetGameByIdResponseDto](
						ctx,
						query,
					)

					Convey("Then it should retrieve game successfully", func() {
						So(result, ShouldNotBeNil)
						So(result.Game, ShouldNotBeNil)
						So(err, ShouldBeNil)

						Convey("And the retrieved game should have the correct ID", func() {
							// Assert that the retrieved game's ID matches the known ID.
							So(result.Game.Id, ShouldEqual, knownGameID.String())
						})

						Convey("And other game properties should be correct", func() {
							// Assert other properties of the retrieved game as needed.
						})
					})
				})
			})
		})

		Convey("Returning a NotFound error when game with specific id does not exist", func() {
			Convey("Given a game with a unknown ID in the database", func() {
				// Create a test context and an unknown game ID.

				query, err := queries.NewGetGameById(unknownGameID)
				So(err, ShouldBeNil)

				Convey("When GetGameById executed for a game with an unknown ID", func() {
					result, err := mediatr.Send[*queries.GetGameById, *dtos.GetGameByIdResponseDto](
						ctx,
						query,
					)

					Convey("Then the game should not be found and null result", func() {
						// Assert that the error indicates that the game was not found.
						So(result, ShouldBeNil)
						So(err, ShouldNotBeNil)
					})
				})
			})
		})

		integrationTestSharedFixture.DisposeTest()
	})
}
