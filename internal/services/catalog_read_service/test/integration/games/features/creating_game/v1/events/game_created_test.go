//go:build integration
// +build integration

package events

// https://github.com/smartystreets/goconvey/wiki

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

	externalEvents "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/creating_game/v1/events/integration_events/external_events"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestGameCreatedConsumer(t *testing.T) {
	// Setup and initialization code here.
	integrationTestSharedFixture := integration.NewIntegrationTestSharedFixture(t)
	// in test mode we set rabbitmq `AutoStart=false` in configuration in rabbitmqOptions, so we should run rabbitmq bus manually
	integrationTestSharedFixture.Bus.Start(context.Background())
	// wait for consumers ready to consume before publishing messages, preparation background workers takes a bit time (for preventing messages lost)
	time.Sleep(1 * time.Second)

	Convey("Game Created Feature", t, func() {
		// will execute with each subtest
		integrationTestSharedFixture.InitializeTest()
		ctx := context.Background()

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey("Consume GameCreated event by consumer", func() {
			fakeGame := &externalEvents.GameCreatedV1{
				Message:     types.NewMessage(uuid.NewV4().String()),
				GameId:      uuid.NewV4().String(),
				Name:        gofakeit.FirstName(),
				Price:       gofakeit.Price(150, 6000),
				CreatedAt:   time.Now(),
				Description: gofakeit.EmojiDescription(),
			}
			hypothesis := messaging.ShouldConsume[*externalEvents.GameCreatedV1](
				ctx,
				integrationTestSharedFixture.Bus,
				nil,
			)

			Convey("When a GameCreated event consumed", func() {
				err := integrationTestSharedFixture.Bus.PublishMessage(ctx, fakeGame, nil)
				So(err, ShouldBeNil)

				Convey("Then it should consume the GameCreated event", func() {
					hypothesis.Validate(ctx, "there is no consumed message", 30*time.Second)
				})
			})
		})

		Convey("Create game in mongo database when a GameCreated event consumed", func() {
			fakeGame := &externalEvents.GameCreatedV1{
				Message:     types.NewMessage(uuid.NewV4().String()),
				GameId:      uuid.NewV4().String(),
				Name:        gofakeit.FirstName(),
				Price:       gofakeit.Price(150, 6000),
				CreatedAt:   time.Now(),
				Description: gofakeit.EmojiDescription(),
			}

			Convey("When a GameCreated event consumed", func() {
				err := integrationTestSharedFixture.Bus.PublishMessage(ctx, fakeGame, nil)
				So(err, ShouldBeNil)

				Convey("It should store game in the mongo database", func() {
					ctx := context.Background()
					pid := uuid.NewV4().String()
					gameCreated := &externalEvents.GameCreatedV1{
						Message:     types.NewMessage(uuid.NewV4().String()),
						GameId:      pid,
						CreatedAt:   time.Now(),
						Name:        gofakeit.Name(),
						Price:       gofakeit.Price(150, 6000),
						Description: gofakeit.AdjectiveDescriptive(),
					}

					err := integrationTestSharedFixture.Bus.PublishMessage(ctx, gameCreated, nil)
					So(err, ShouldBeNil)

					var game *models.Game

					err = testUtils.WaitUntilConditionMet(func() bool {
						game, err = integrationTestSharedFixture.GameRepository.GetGameByGameId(ctx, pid)

						return err == nil && game != nil
					})

					So(err, ShouldBeNil)
					So(game, ShouldNotBeNil)
					So(game.GameId, ShouldEqual, pid)
				})
			})
		})

		integrationTestSharedFixture.DisposeTest()
	})

	integrationTestSharedFixture.Log.Info("TearDownSuite started")
	integrationTestSharedFixture.Bus.Stop()
	time.Sleep(1 * time.Second)
}
