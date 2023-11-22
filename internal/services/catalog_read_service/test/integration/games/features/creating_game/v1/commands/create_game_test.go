//go:build integration
// +build integration

package commands

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mehdihadeli/go-mediatr"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/creating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/creating_game/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestCreateGame(t *testing.T) {
	integrationTestSharedFixture := integration.NewIntegrationTestSharedFixture(t)

	Convey("Creating Game Feature", t, func() {
		ctx := context.Background()
		integrationTestSharedFixture.InitializeTest()

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey(
			"Creating a new game and saving it to the database for a none-existing game",
			func() {
				Convey("Given new game doesn't exists in the system", func() {
					command, err := commands.NewCreateGame(
						uuid.NewV4().String(),
						gofakeit.Name(),
						gofakeit.AdjectiveDescriptive(),
						gofakeit.Price(150, 6000),
						time.Now(),
					)
					So(err, ShouldBeNil)

					Convey(
						"When the CreateGame command is executed and game doesn't exists",
						func() {
							result, err := mediatr.Send[*commands.CreateGame, *dtos.CreateGameResponseDto](
								ctx,
								command,
							)

							Convey("Then the game should be created successfully", func() {
								So(err, ShouldBeNil)
								So(result, ShouldNotBeNil)

								Convey(
									"And the game ID should not be empty and same as commandId",
									func() {
										So(result.Id, ShouldEqual, command.Id)

										Convey(
											"And game detail should be retrievable from the database",
											func() {
												createdGame, err := integrationTestSharedFixture.GameRepository.GetGameById(
													ctx,
													result.Id,
												)
												So(err, ShouldBeNil)
												So(createdGame, ShouldNotBeNil)
											},
										)
									},
								)
							})
						},
					)
				})
			},
		)

		integrationTestSharedFixture.DisposeTest()
	})
}
