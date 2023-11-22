//go:build e2e
// +build e2e

package v1

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestCreateGameEndpoint(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "CreateGame Endpoint EndToEnd Tests")
}

var _ = Describe("CreateGame Feature", func() {
	var (
		ctx     context.Context
		request *dtos.CreateGameRequestDto
	)

	_ = BeforeEach(func() {
		ctx = context.Background()

		By("Seeding the required data")
		integrationFixture.InitializeTest()
	})

	_ = AfterEach(func() {
		By("Cleanup test data")
		integrationFixture.DisposeTest()
	})

	// "Scenario" step for testing the create game API with valid input
	Describe("Create new game return created status with valid input", func() {
		BeforeEach(func() {
			// Generate a valid request
			request = &dtos.CreateGameRequestDto{
				Description: gofakeit.AdjectiveDescriptive(),
				Price:       gofakeit.Price(100, 1000),
				Name:        gofakeit.Name(),
			}
		})
		// "When" step
		When("A valid request is made to create a game", func() {
			// "Then" step
			It("Should returns a StatusCreated response", func() {
				// Create an HTTPExpect instance and make the request
				expect := httpexpect.New(GinkgoT(), integrationFixture.BaseAddress)
				expect.POST("games").
					WithContext(ctx).
					WithJSON(request).
					Expect().
					Status(http.StatusCreated)
			})
		})
	})

	// "Scenario" step for testing the create game API with invalid price input
	Describe("Create game returns a BadRequest status with invalid price input", func() {
		BeforeEach(func() {
			// Generate an invalid request with zero price
			request = &dtos.CreateGameRequestDto{
				Description: gofakeit.AdjectiveDescriptive(),
				Price:       0.0,
				Name:        gofakeit.Name(),
			}
		})
		// "When" step
		When("An invalid request is made with a zero price", func() {
			// "Then" step
			It("Should return a BadRequest status", func() {
				// Create an HTTPExpect instance and make the request
				expect := httpexpect.New(GinkgoT(), integrationFixture.BaseAddress)
				expect.POST("games").
					WithContext(ctx).
					WithJSON(request).
					Expect().
					Status(http.StatusBadRequest)
			})
		})
	})
})
