//go:build integration
// +build integration

package v1

import (
	"context"
	"fmt"
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

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/events/integration_events"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestUpdateGame(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "Updated Games Integration Tests")
}

var _ = Describe("Update Game Feature", func() {
	// Define variables to hold command and result data
	var (
		ctx           context.Context
		existingGame  *models.Game
		command       *commands.UpdateGame
		result        *mediatr.Unit
		err           error
		id            uuid.UUID
		shouldPublish hypothesis.Hypothesis[*integration_events.GameUpdatedV1]
	)

	_ = BeforeEach(func() {
		By("Seeding the required data")
		integrationFixture.InitializeTest()

		existingGame = integrationFixture.Items[0]
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

	// "Scenario" step for testing updating an existing game
	Describe("Updating an existing game in the database", func() {
		Context("Given game exists in the database", func() {
			BeforeEach(func() {
				command, err = commands.NewUpdateGame(
					existingGame.GameId,
					"Updated Game Name",
					existingGame.Description,
					existingGame.Price,
				)
				Expect(err).NotTo(HaveOccurred())
			})

			// "When" step
			When("the UpdateGame command is executed", func() {
				BeforeEach(func() {
					result, err = mediatr.Send[*commands.UpdateGame, *mediatr.Unit](ctx, command)
				})

				// "Then" step
				It("Should not return an error", func() {
					Expect(err).NotTo(HaveOccurred())
					Expect(result).NotTo(BeNil())
				})

				It("Should return a non-nil result", func() {
					Expect(result).NotTo(BeNil())
				})

				It("Should update the existing game in the database", func() {
					updatedGame, err := integrationFixture.GameRepository.GetGameById(
						ctx,
						existingGame.GameId,
					)
					Expect(err).To(BeNil())
					Expect(updatedGame).NotTo(BeNil())
					Expect(updatedGame.GameId).To(Equal(existingGame.GameId))
					Expect(updatedGame.Price).To(Equal(existingGame.Price))
					Expect(updatedGame.Name).NotTo(Equal(existingGame.Name))
				})
			})
		})
	})

	// "Scenario" step for testing updating a non-existing game
	Describe("Updating a non-existing game in the database", func() {
		Context("Given game not exists in the database", func() {
			BeforeEach(func() {
				// Generate a random ID that does not exist in the database
				id = uuid.NewV4()
				command, err = commands.NewUpdateGame(
					id,
					"Updated Game Name",
					"Updated Game Description",
					100,
				)
				Expect(err).NotTo(HaveOccurred())
			})

			// "When" step
			When("the UpdateGame command executed for non-existing game", func() {
				BeforeEach(func() {
					result, err = mediatr.Send[*commands.UpdateGame, *mediatr.Unit](ctx, command)
				})

				// "Then" step
				It("Should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
				It("Should not return a result", func() {
					Expect(result).To(BeNil())
				})

				It("Should return a NotFound error", func() {
					Expect(err).To(MatchError(ContainSubstring(fmt.Sprintf("game with id %s not found", id))))
				})

				It("Should return a custom NotFound error", func() {
					Expect(customErrors.IsNotFoundError(err)).To(BeTrue())
					Expect(customErrors.IsApplicationError(err, http.StatusNotFound)).To(BeTrue())
				})
			})
		})
	})

	// "Scenario" step for testing updating an existing game
	Describe("Publishing GameUpdated when game updated  successfully", func() {
		Context("Given game exists in the database", func() {
			BeforeEach(func() {
				command, err = commands.NewUpdateGame(
					existingGame.GameId,
					"Updated Game Name",
					existingGame.Description,
					existingGame.Price,
				)
				Expect(err).NotTo(HaveOccurred())

				shouldPublish = messaging.ShouldProduced[*integration_events.GameUpdatedV1](
					ctx,
					integrationFixture.Bus,
					nil,
				)
			})

			// "When" step
			When("the UpdateGame command is executed for existing game", func() {
				BeforeEach(func() {
					result, err = mediatr.Send[*commands.UpdateGame, *mediatr.Unit](ctx, command)
				})

				It("Should publish GameUpdated event to the broker", func() {
					// ensuring message published to the rabbitmq broker
					shouldPublish.Validate(ctx, "there is no published message", time.Second*30)
				})
			})
		})
	})
})
