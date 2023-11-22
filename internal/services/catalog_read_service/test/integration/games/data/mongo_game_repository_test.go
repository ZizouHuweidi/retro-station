//go:build integration
// +build integration

package data

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestGamePostgresRepository(t *testing.T) {
	integrationTestSharedFixture := integration.NewIntegrationTestSharedFixture(t)

	// scenario
	Convey("Game Repository", t, func() {
		integrationTestSharedFixture.InitializeTest()

		Convey("When we create the new game in the database", func() {
			ctx := context.Background()
			game := &models.Game{
				Id:          uuid.NewV4().String(),
				GameId:      uuid.NewV4().String(),
				Name:        gofakeit.Name(),
				Description: gofakeit.AdjectiveDescriptive(),
				Price:       gofakeit.Price(100, 1000),
				CreatedAt:   time.Now(),
			}

			createdGame, err := integrationTestSharedFixture.GameRepository.CreateGame(ctx, game)

			Convey("Then the game should be created successfully", func() {
				// Assert that there is no error during creation.
				So(err, ShouldBeNil)

				Convey("And we should be able to retrieve the game by ID", func() {
					retrievedGame, err := integrationTestSharedFixture.GameRepository.GetGameById(
						ctx,
						createdGame.Id,
					)

					Convey("And retrieved game should match the created game", func() {
						// Assert that there is no error during retrieval.
						So(err, ShouldBeNil)

						// Assert that the retrieved game matches the created game.
						So(retrievedGame.Id, ShouldEqual, createdGame.Id)
					})
				})
			})
		})

		Convey("When we delete the existing game", func() {
			ctx := context.Background()

			id := integrationTestSharedFixture.Items[0].Id
			err := integrationTestSharedFixture.GameRepository.DeleteGameByID(ctx, id)

			Convey("Then the game should be deleted successfully", func() {
				// Ensure there is no error during deletion.
				So(err, ShouldBeNil)

				Convey("And when we attempt to retrieve the game by ID", func() {
					game, err := integrationTestSharedFixture.GameRepository.GetGameById(ctx, id)

					Convey("And error should occur indicating the game is not found", func() {
						// Verify that there is an error.
						So(err, ShouldNotBeNil)

						// Check if the error is of a specific type (e.g., a not found error).
						So(customErrors.IsNotFoundError(err), ShouldBeTrue)

						// Verify that the retrieved game is nil.
						So(game, ShouldBeNil)
					})
				})
			})
		})

		Convey("When we update the existing game", func() {
			Convey("Then the game should be updated successfully", func() {
				ctx := context.Background()

				id := integrationTestSharedFixture.Items[0].Id
				existingGame, err := integrationTestSharedFixture.GameRepository.GetGameById(ctx, id)

				// Make sure the existing game exists and there is no error.
				So(err, ShouldBeNil)
				So(existingGame, ShouldNotBeNil)

				// Modify the existing game's name.
				existingGame.Name = "test_update_game"

				// Update the game in the database.
				_, err = integrationTestSharedFixture.GameRepository.UpdateGame(ctx, existingGame)

				// Ensure there is no error during the update.
				So(err, ShouldBeNil)

				// Retrieve the updated game from the database.
				updatedGame, err := integrationTestSharedFixture.GameRepository.GetGameById(ctx, id)
				So(err, ShouldBeNil)

				// Verify that the updated game's name matches the new name.
				So(updatedGame.Name, ShouldEqual, "test_update_game")
			})
		})

		Convey("When attempting to get a game that does not exist", func() {
			ctx := context.Background()

			res, err := integrationTestSharedFixture.GameRepository.GetGameById(ctx, uuid.NewV4().String())

			Convey("Then it should return a NotFound error and nil result", func() {
				// Verify that there is an error.
				So(err, ShouldNotBeNil)

				// Check if the error is of a specific type (e.g., a not found error).
				So(customErrors.IsNotFoundError(err), ShouldBeTrue)

				// Verify that the retrieved result is nil.
				So(res, ShouldBeNil)
			})
		})

		Convey("When attempting to get an existing game from the database", func() {
			ctx := context.Background()

			id := integrationTestSharedFixture.Items[0].Id
			res, err := integrationTestSharedFixture.GameRepository.GetGameById(ctx, id)

			Convey("Then it should return the game and no error", func() {
				// Ensure there is no error.
				So(err, ShouldBeNil)

				// Verify that the result is not nil.
				So(res, ShouldNotBeNil)

				// Verify that the retrieved game's ID matches the expected ID.
				So(res.Id, ShouldEqual, id)
			})
		})

		Convey("When attempting to get all existing games from the database", func() {
			ctx := context.Background()

			res, err := integrationTestSharedFixture.GameRepository.GetAllGames(ctx, utils.NewListQuery(10, 1))

			Convey("Then it should return the list of games and no error", func() {
				// Ensure there is no error.
				So(err, ShouldBeNil)

				// Verify the expected number of games in the list.
				So(len(res.Items), ShouldEqual, 2)
			})
		})

		integrationTestSharedFixture.DisposeTest()
	})
}
