//go:build integration
// +build integration

package events

import (
	"context"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/messaging"
	testUtils "github.com/zizouhuweidi/retro-station/internal/pkg/test/utils"

	externalEvents "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/deleting_games/v1/events/integration_events/external_events"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestGameDeleted(t *testing.T) {
	// Setup and initialization code here.
	integrationTestSharedFixture := integration.NewIntegrationTestSharedFixture(t)
	// in test mode we set rabbitmq `AutoStart=false` in configuration in rabbitmqOptions, so we should run rabbitmq bus manually
	integrationTestSharedFixture.Bus.Start(context.Background())
	// wait for consumers ready to consume before publishing messages, preparation background workers takes a bit time (for preventing messages lost)
	time.Sleep(1 * time.Second)

	Convey("Game Deleted Feature", t, func() {
		ctx := context.Background()
		// will execute with each subtest
		integrationTestSharedFixture.InitializeTest()

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey("Consume GameDeleted event by consumer", func() {
			event := &externalEvents.GameDeletedV1{
				Message: types.NewMessage(uuid.NewV4().String()),
				GameId:  integrationTestSharedFixture.Items[0].GameId,
			}
			// check for consuming `GameDeletedV1` message with existing consumer
			hypothesis := messaging.ShouldConsume[*externalEvents.GameDeletedV1](
				ctx,
				integrationTestSharedFixture.Bus,
				nil,
			)

			Convey("When a GameDeleted event consumed", func() {
				err := integrationTestSharedFixture.Bus.PublishMessage(
					ctx,
					event,
					nil,
				)
				So(err, ShouldBeNil)

				Convey("Then it should consume the GameDeleted event", func() {
					hypothesis.Validate(ctx, "there is no consumed message", 30*time.Second)
				})
			})
		})

		// https://specflow.org/learn/gherkin/#learn-gherkin
		// scenario
		Convey("Delete game in mongo database when a GameDeleted event consumed", func() {
			event := &externalEvents.GameDeletedV1{
				Message: types.NewMessage(uuid.NewV4().String()),
				GameId:  integrationTestSharedFixture.Items[0].GameId,
			}

			Convey("When a GameDeleted event consumed", func() {
				err := integrationTestSharedFixture.Bus.PublishMessage(
					ctx,
					event,
					nil,
				)
				So(err, ShouldBeNil)

				Convey("It should delete game in the mongo database", func() {
					ctx := context.Background()

					gameDeleted := &externalEvents.GameDeletedV1{
						Message: types.NewMessage(uuid.NewV4().String()),
						GameId:  integrationTestSharedFixture.Items[0].GameId,
					}

					err := integrationTestSharedFixture.Bus.PublishMessage(ctx, gameDeleted, nil)
					So(err, ShouldBeNil)

					var p *models.Game

					So(testUtils.WaitUntilConditionMet(func() bool {
						p, err = integrationTestSharedFixture.GameRepository.GetGameByGameId(
							ctx,
							integrationTestSharedFixture.Items[0].GameId,
						)
						So(err, ShouldBeNil)

						return p == nil
					}), ShouldBeNil)

					So(p, ShouldBeNil)
				})
			})
		})

		integrationTestSharedFixture.DisposeTest()
	})

	integrationTestSharedFixture.Log.Info("TearDownSuite started")
	integrationTestSharedFixture.Bus.Stop()
	time.Sleep(1 * time.Second)
}
