//go:build integration
// +build integration

package v1

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mehdihadeli/go-mediatr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/hypothesis"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/messaging"

	createGameCommand "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/dtos"
	integrationEvents "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/events/integration_events"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestCreateGame(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "Create Game Integration Tests")
}

// https://specflow.org/learn/gherkin/#learn-gherkin
// scenario
var _ = Describe("Creating Game Feature", func() {
	var (
		ctx           context.Context
		err           error
		command       *createGameCommand.CreateGame
		result        *dtos.CreateGameResponseDto
		createdGame   *models.Game
		id            uuid.UUID
		shouldPublish hypothesis.Hypothesis[*integrationEvents.GameCreatedV1]
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

	// "Scenario" step for testing creating a new game
	Describe("Creating a new game and saving it to the database when game doesn't exists", func() {
		Context("Given new game doesn't exists in the system", func() {
			BeforeEach(func() {
				command, err = createGameCommand.NewCreateGame(
					gofakeit.Name(),
					gofakeit.AdjectiveDescriptive(),
					gofakeit.Price(150, 6000),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(command).ToNot(BeNil())
			})

			When("the CreateGame command is executed for non-existing game", func() {
				BeforeEach(func() {
					result, err = mediatr.Send[*createGameCommand.CreateGame, *dtos.CreateGameResponseDto](
						ctx,
						command,
					)
				})

				It("Should create the game successfully", func() {
					Expect(err).NotTo(HaveOccurred())
					Expect(result).NotTo(BeNil())
				})

				It("Should have a non-empty game ID matching the command ID", func() {
					Expect(result.GameID).To(Equal(command.GameID))
				})

				It("Should be able to retrieve the game from the database", func() {
					createdGame, err = integrationFixture.GameRepository.GetGameById(
						ctx,
						result.GameID,
					)
					Expect(err).NotTo(HaveOccurred())

					Expect(result).NotTo(BeNil())
					Expect(command.GameID).To(Equal(result.GameID))
					Expect(createdGame).NotTo(BeNil())
				})
			})
		})
	})

	// "Scenario" step for testing creating a game with duplicate data
	Describe("Creating a new game with duplicate data and already exists game", func() {
		Context("Given game already exists in the system", func() {
			BeforeEach(func() {
				command = &createGameCommand.CreateGame{
					Name:        gofakeit.Name(),
					Description: gofakeit.AdjectiveDescriptive(),
					Price:       gofakeit.Price(150, 6000),
					GameID:      id,
				}
			})

			When("the CreateGame command is executed for existing game", func() {
				BeforeEach(func() {
					result, err = mediatr.Send[*createGameCommand.CreateGame, *dtos.CreateGameResponseDto](
						ctx,
						command,
					)
				})

				It("Should return an error indicating duplicate record", func() {
					Expect(err).To(HaveOccurred())
					Expect(customErrors.IsApplicationError(err, http.StatusConflict)).To(BeTrue())
				})

				It("Should not return a result", func() {
					Expect(result).To(BeNil())
				})
			})
		})
	})

	// "Scenario" step for testing creating a game with duplicate data
	Describe("Publishing GameCreated event to the broker when game saved successfully", func() {
		Context("Given new game doesn't exists in the system", func() {
			BeforeEach(func() {
				shouldPublish = messaging.ShouldProduced[*integrationEvents.GameCreatedV1](
					ctx,
					integrationFixture.Bus,
					nil,
				)
				command, err = createGameCommand.NewCreateGame(
					gofakeit.Name(),
					gofakeit.AdjectiveDescriptive(),
					gofakeit.Price(150, 6000),
				)
				Expect(err).ToNot(HaveOccurred())
			})

			When("CreateGame command is executed for non-existing game", func() {
				BeforeEach(func() {
					result, err = mediatr.Send[*createGameCommand.CreateGame, *dtos.CreateGameResponseDto](
						ctx,
						command,
					)
				})

				It("Should return no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should return not nil result", func() {
					Expect(result).ToNot(BeNil())
				})

				It("Should publish GameCreated event to the broker", func() {
					// ensuring message published to the rabbitmq broker
					shouldPublish.Validate(ctx, "there is no published message", time.Second*30)
				})
			})
		})
	})
})
