//go:build e2e
// +build e2e

package grpc

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"

	gameService "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc/genproto"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestGameGrpcServiceEndToEnd(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "GameGrpcService EndToEnd Tests")
}

var _ = Describe("Game Grpc Service Feature", func() {
	var (
		ctx context.Context
		id  uuid.UUID
	)

	_ = BeforeEach(func() {
		ctx = context.Background()

		By("Seeding the required data")
		integrationFixture.InitializeTest()

		id = integrationFixture.Items[0].GameId
	})

	_ = AfterEach(func() {
		By("Cleanup test data")
		integrationFixture.DisposeTest()
	})

	// "Scenario" step for testing the creation of a game with valid data in the database
	Describe("Creation of a game with valid data in the database", func() {
		// "When" step
		When("A request is made to create a game with valid data", func() {
			// "Then" step
			It("Should return a non-empty GameId", func() {
				// Create a gRPC request with valid data
				request := &gameService.CreateGameReq{
					Price:       gofakeit.Price(100, 1000),
					Name:        gofakeit.Name(),
					Description: gofakeit.AdjectiveDescriptive(),
				}

				// Make the gRPC request to create the game
				res, err := integrationFixture.GameServiceClient.CreateGame(ctx, request)
				Expect(err).To(BeNil())
				Expect(res).NotTo(BeNil())
				Expect(res.GameId).NotTo(BeEmpty())
			})
		})
	})

	// "Scenario" step for testing the retrieval of data with a valid ID
	Describe("Retrieve game with a valid ID", func() {
		// "When" step
		When("A request is made to retrieve data with a valid ID", func() {
			// "Then" step
			It("Should return data with a matching GameId", func() {
				// Make the gRPC request to retrieve data by ID
				res, err := integrationFixture.GameServiceClient.GetGameById(
					ctx,
					&gameService.GetGameByIdReq{GameId: id.String()},
				)

				Expect(err).To(BeNil())
				Expect(res).NotTo(BeNil())
				Expect(res.Game).NotTo(BeNil())
				Expect(res.Game.GameId).To(Equal(id.String()))
			})
		})
	})
})
