//go:build integration
// +build integration

package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestGamePostgresRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "GamePostgresRepository Integration Tests")
}

var _ = Describe("Game Repository Suite", func() {
	// Define variables to hold repository and game data
	var (
		ctx          context.Context
		game         *models.Game
		createdGame  *models.Game
		updatedGame  *models.Game
		existingGame *models.Game
		err          error
		id           uuid.UUID
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

	// "Scenario" step for testing creating a new game in the database
	Describe("Creating a new game in the database", func() {
		BeforeEach(func() {
			game = &models.Game{
				Name:        gofakeit.Name(),
				Description: gofakeit.AdjectiveDescriptive(),
				GameId:      uuid.NewV4(),
				Price:       gofakeit.Price(100, 1000),
				CreatedAt:   time.Now(),
			}
		})

		// "When" step
		When("CreateGame function of GameRepository executed", func() {
			BeforeEach(func() {
				createdGame, err = integrationFixture.GameRepository.CreateGame(ctx, game)
			})

			// "Then" step
			It("Should not return an error", func() {
				Expect(err).To(BeNil())
			})

			It("Should return a non-nil created game", func() {
				Expect(createdGame).NotTo(BeNil())
			})

			It("Should have the same GameId as the input game", func() {
				Expect(createdGame.GameId).To(Equal(game.GameId))
			})

			It("Should be able to retrieve the created game from the database", func() {
				retrievedGame, err := integrationFixture.GameRepository.GetGameById(
					ctx,
					createdGame.GameId,
				)
				Expect(err).NotTo(HaveOccurred())
				Expect(retrievedGame).NotTo(BeNil())
				Expect(retrievedGame.GameId).To(Equal(createdGame.GameId))
			})
		})
	})

	// "Scenario" step for testing updating an existing game in the database
	Describe("Updating an existing game in the database", func() {
		BeforeEach(func() {
			existingGame, err = integrationFixture.GameRepository.GetGameById(ctx, id)
			Expect(err).To(BeNil())
			Expect(existingGame).NotTo(BeNil())
		})

		// "When" step
		When("UpdateGame function of GameRepository executed", func() {
			BeforeEach(func() {
				// Update the name of the existing game
				existingGame.Name = "Updated Game Name"
				_, err = integrationFixture.GameRepository.UpdateGame(ctx, existingGame)
			})

			// "Then" step
			It("Should not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("Should be able to retrieve the updated game from the database", func() {
				updatedGame, err = integrationFixture.GameRepository.GetGameById(
					ctx,
					existingGame.GameId,
				)
				Expect(err).To(BeNil())
				Expect(updatedGame).NotTo(BeNil())
				Expect(updatedGame.Name).To(Equal("Updated Game Name"))
				// You can add more assertions to validate other properties of the updated game
			})
		})
	})

	// "Scenario" step for testing deleting an existing game in the database
	Describe("Deleting an existing game from the database", func() {
		BeforeEach(func() {
			// Ensure that the game with 'id' exists in the database
			game, err := integrationFixture.GameRepository.GetGameById(ctx, id)
			Expect(err).To(BeNil())
			Expect(game).NotTo(BeNil())
		})

		// "When" step
		When("DeleteGame function of GameRepository executed", func() {
			BeforeEach(func() {
				err = integrationFixture.GameRepository.DeleteGameByID(ctx, id)
			})

			// "Then" step
			It("Should not return an error", func() {
				Expect(err).To(BeNil())
			})

			It("Should delete given game from the database", func() {
				game, err := integrationFixture.GameRepository.GetGameById(ctx, id)
				Expect(err).To(HaveOccurred())
				Expect(customErrors.IsNotFoundError(err)).To(BeTrue())
				Expect(game).To(BeNil())
			})
		})
	})

	// "Scenario" step for testing retrieving an existing game from the database
	Describe("Retrieving an existing game from the database", func() {
		BeforeEach(func() {
			// Ensure that the game with 'id' exists in the database
			game, err := integrationFixture.GameRepository.GetGameById(ctx, id)
			Expect(err).To(BeNil())
			Expect(game).NotTo(BeNil())
		})

		// "When" step
		When("GetGameById function of GameRepository executed", func() {
			BeforeEach(func() {
				existingGame, err = integrationFixture.GameRepository.GetGameById(ctx, id)
			})
			It("should not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
				Expect(existingGame).NotTo(BeNil())
			})

			It("should retrieve correct data from database by Id", func() {
				Expect(existingGame.GameId).To(Equal(id))
			})
		})
	})

	// "Scenario" step for testing retrieving a game that does not exist in the database
	Describe("Retrieving a game that does not exist in the database", func() {
		BeforeEach(func() {
			// Ensure that the game with 'id' exists in the database
			game, err := integrationFixture.GameRepository.GetGameById(ctx, id)
			Expect(err).To(BeNil())
			Expect(game).NotTo(BeNil())
		})

		// "When" step
		When("GetGameById function of GameRepository executed", func() {
			BeforeEach(func() {
				// Use a random UUID that does not exist in the database
				nonexistentID := uuid.NewV4()
				existingGame, err = integrationFixture.GameRepository.GetGameById(ctx, nonexistentID)
			})

			// "Then" step
			It("Should return a NotFound error", func() {
				Expect(err).To(HaveOccurred())
				Expect(customErrors.IsNotFoundError(err)).To(BeTrue())
			})

			It("Should not return a game", func() {
				Expect(existingGame).To(BeNil())
			})
		})
	})

	// "Scenario" step for testing retrieving all existing games from the database
	Describe("Retrieving all existing games from the database", func() {
		// "When" step
		When("GetAllGames function of GameRepository executed", func() {
			It("should not return an error and return the correct number of games", func() {
				res, err := integrationFixture.GameRepository.GetAllGames(ctx, utils.NewListQuery(10, 1))
				Expect(err).To(BeNil())
				Expect(res).NotTo(BeNil())
				Expect(len(res.Items)).To(Equal(2)) // Replace with the expected number of games
			})
		})
	})
})
