//go:build integration
// +build integration

package events

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/messaging"
	testUtils "github.com/zizouhuweidi/retro-station/internal/pkg/test/utils"

	externalEvents "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/updating_games/v1/events/integration_events/external_events"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestGameUpdatedConsumer(t *testing.T) {
	// Setup and initialization code here.
	integrationTestSharedFixture := integration.NewIntegrationTestSharedFixture(
		t,
	)
	// in test mode we set rabbitmq `AutoStart=false` in configuration in rabbitmqOptions, so we should run rabbitmq bus manually
	integrationTestSharedFixture.Bus.Start(context.Background())
	// wait for consumers ready to consume before publishing messages, preparation background workers takes a bit time (for preventing messages lost)
	time.Sleep(1 * time.Second)

	Convey("Game Created Feature", t, func() {
		ctx := context.Background()
		integrationTestSharedFixture.InitializeTest()

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey("Consume GameUpdated event by consumer", func() {
			hypothesis := messaging.ShouldConsume[*externalEvents.GameUpdatedV1](
				ctx,
				integrationTestSharedFixture.Bus,
				nil,
			)

			fakeUpdateGame := &externalEvents.GameUpdatedV1{
				Message:     types.NewMessage(uuid.NewV4().String()),
				GameId:      integrationTestSharedFixture.Items[0].GameId,
				Name:        gofakeit.Name(),
				Price:       gofakeit.Price(100, 1000),
				Description: gofakeit.EmojiDescription(),
				UpdatedAt:   time.Now(),
			}

			Convey("When a GameUpdated event consumed", func() {
				err := integrationTestSharedFixture.Bus.PublishMessage(
					ctx,
					fakeUpdateGame,
					nil,
				)
				So(err, ShouldBeNil)

				Convey(
					"Then it should consume the GameUpdated event",
					func() {
						hypothesis.Validate(
							ctx,
							"there is no consumed message",
							30*time.Second,
						)
					},
				)
			})
		})

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey(
			"Update game in mongo database when a GameDeleted event consumed",
			func() {
				fakeUpdateGame := &externalEvents.GameUpdatedV1{
					Message:     types.NewMessage(uuid.NewV4().String()),
					GameId:      integrationTestSharedFixture.Items[0].GameId,
					Name:        gofakeit.Name(),
					Price:       gofakeit.Price(100, 1000),
					Description: gofakeit.EmojiDescription(),
					UpdatedAt:   time.Now(),
				}

				Convey("When a GameUpdated event consumed", func() {
					err := integrationTestSharedFixture.Bus.PublishMessage(
						ctx,
						fakeUpdateGame,
						nil,
					)
					So(err, ShouldBeNil)

					Convey(
						"Then It should update game in the mongo database",
						func() {
							ctx := context.Background()
							gameUpdated := &externalEvents.GameUpdatedV1{
								Message: types.NewMessage(
									uuid.NewV4().String(),
								),
								GameId:      integrationTestSharedFixture.Items[0].GameId,
								Name:        gofakeit.Name(),
								Description: gofakeit.AdjectiveDescriptive(),
								Price:       gofakeit.Price(150, 6000),
								UpdatedAt:   time.Now(),
							}

							err := integrationTestSharedFixture.Bus.PublishMessage(
								ctx,
								gameUpdated,
								nil,
							)
							So(err, ShouldBeNil)

							var game *models.Game

							err = testUtils.WaitUntilConditionMet(func() bool {
								game, err = integrationTestSharedFixture.GameRepository.GetGameByGameId(
									ctx,
									integrationTestSharedFixture.Items[0].GameId,
								)

								return game != nil &&
									game.Name == gameUpdated.Name
							})

							So(err, ShouldBeNil)
							So(game, ShouldNotBeNil)
							So(
								gameUpdated.GameId,
								ShouldEqual,
								game.GameId,
							)
						},
					)
				})
			},
		)

		integrationTestSharedFixture.DisposeTest()
	})

	integrationTestSharedFixture.Log.Info("TearDownSuite started")
	integrationTestSharedFixture.Bus.Stop()
	time.Sleep(1 * time.Second)
}
