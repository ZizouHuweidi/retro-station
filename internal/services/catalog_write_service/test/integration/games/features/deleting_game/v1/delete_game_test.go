//go:build integration
// +build integration

package v1

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/mehdihadeli/go-mediatr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/hypothesis"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/messaging"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/deleting_game/v1/commands"
	integrationEvents "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/deleting_game/v1/events/integration_events"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestDeleteGame(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "Delete Game Integration Tests")
}

// https://specflow.org/learn/gherkin/#learn-gherkin
// scenario
var _ = Describe("Delete Game Feature", func() {
	var (
		ctx           context.Context
		err           error
		command       *commands.DeleteGame
		result        *mediatr.Unit
		id            uuid.UUID
		notExistsId   uuid.UUID
		shouldPublish hypothesis.Hypothesis[*integrationEvents.GameDeletedV1]
	)

	_ = BeforeEach(func() {
		By("Seeding the required data")
		integrationFixture.InitializeTest()

		id = integrationFixture.Items[0].GameId
	})

	_ = AfterEach(func() {
		By("Cleanup test data")
		integrationFixture.DisposeTest()
	})

	_ = BeforeSuite(func() {
		ctx = context.Background()

		// in test mode we set rabbitmq `AutoStart=false` in configuration in rabbitmqOptions, so we should run rabbitmq bus manually
		err = integrationFixture.Bus.Start(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		// wait for consumers ready to consume before publishing messages, preparation background workers takes a bit time (for preventing messages lost)
		time.Sleep(1 * time.Second)
	})

	_ = AfterSuite(func() {
		integrationFixture.Log.Info("TearDownSuite started")
		err := integrationFixture.Bus.Stop()
		Expect(err).ShouldNot(HaveOccurred())
		time.Sleep(1 * time.Second)
	})

	// "Scenario" step for testing deleting an existing game
	Describe("Deleting an existing game from the database", func() {
		Context("Given game already exists in the system", func() {
			BeforeEach(func() {
				shouldPublish = messaging.ShouldProduced[*integrationEvents.GameDeletedV1](
					ctx,
					integrationFixture.Bus,
					nil,
				)
				command, err = commands.NewDeleteGame(id)
				Expect(err).ShouldNot(HaveOccurred())
			})

			When("the DeleteGame command is executed for existing game", func() {
				BeforeEach(func() {
					result, err = mediatr.Send[*commands.DeleteGame, *mediatr.Unit](
						ctx,
						command,
					)
				})

				It("Should not return an error", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("Should delete the game from the database", func() {
					deletedGame, err := integrationFixture.GameRepository.GetGameById(ctx, id)
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError(ContainSubstring("can't find the game with id")))
					Expect(deletedGame).To(BeNil())
				})
			})
		})
	})

	// "Scenario" step for testing deleting a non-existing game
	Describe("Deleting a non-existing game from the database", func() {
		Context("Given game does not exists in the system", func() {
			BeforeEach(func() {
				notExistsId = uuid.NewV4()
				command, err = commands.NewDeleteGame(notExistsId)
				Expect(err).ShouldNot(HaveOccurred())
			})

			When("the DeleteGame command is executed for non-existing game", func() {
				BeforeEach(func() {
					result, err = mediatr.Send[*commands.DeleteGame, *mediatr.Unit](
						ctx,
						command,
					)
				})

				It("Should return an error", func() {
					Expect(err).To(HaveOccurred())
				})

				It("Should return a NotFound error", func() {
					Expect(err).To(MatchError(ContainSubstring("game not found")))
				})

				It("Should return a custom NotFound error", func() {
					Expect(customErrors.IsApplicationError(err, http.StatusNotFound)).To(BeTrue())
					Expect(customErrors.IsNotFoundError(err)).To(BeTrue())
				})

				It("Should not return a result", func() {
					Expect(result).To(BeNil())
				})
			})
		})
	})

	Describe("Publishing GameDeleted event when game deleted successfully", func() {
		Context("Given game already exists in the system", func() {
			BeforeEach(func() {
				shouldPublish = messaging.ShouldProduced[*integrationEvents.GameDeletedV1](
					ctx,
					integrationFixture.Bus,
					nil,
				)
				command, err = commands.NewDeleteGame(id)
				Expect(err).ShouldNot(HaveOccurred())
			})

			When("the DeleteGame command is executed for existing game", func() {
				BeforeEach(func() {
					result, err = mediatr.Send[*commands.DeleteGame, *mediatr.Unit](
						ctx,
						command,
					)
				})

				It("Should publish GameDeleted event to the broker", func() {
					// ensuring message published to the rabbitmq broker
					shouldPublish.Validate(ctx, "there is no published message", time.Second*30)
				})
			})
		})
	})
})
