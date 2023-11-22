//go:build integration
// +build integration

package queries

import (
	"context"
	"testing"

	"github.com/mehdihadeli/go-mediatr"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/getting_games/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/getting_games/v1/queries"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestGetGames(t *testing.T) {
	integrationTestSharedFixture := integration.NewIntegrationTestSharedFixture(t)

	Convey("Get All Games Feature", t, func() {
		ctx := context.Background()
		integrationTestSharedFixture.InitializeTest()

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey("Getting a list of existing games from the database", func() {
			Convey("Given a set of existing games in the system", func() {
				query := queries.NewGetGames(utils.NewListQuery(10, 1))

				Convey("When GetGame query executed for existing games", func() {
					queryResult, err := mediatr.Send[*queries.GetGames, *dtos.GetGamesResponseDto](
						ctx,
						query,
					)

					Convey("Then the games should be retrieved successfully", func() {
						// Assert that the error is nil (indicating success).
						So(err, ShouldBeNil)
						So(queryResult, ShouldNotBeNil)
						So(queryResult.Games, ShouldNotBeNil)

						Convey("And the list of games should not be empty", func() {
							// Assert that the list of games is not empty.
							So(queryResult.Games.Items, ShouldNotBeEmpty)

							Convey("And each game should have the correct properties", func() {
								for _, game := range queryResult.Games.Items {
									// Assert properties of each game as needed.
									// For example:
									So(game.Name, ShouldNotBeEmpty)
									So(game.Price, ShouldBeGreaterThan, 0.0)
									// Add more assertions as needed.
								}
							})
						})
					})
				})
			})
		})

		integrationTestSharedFixture.DisposeTest()
	})
}
