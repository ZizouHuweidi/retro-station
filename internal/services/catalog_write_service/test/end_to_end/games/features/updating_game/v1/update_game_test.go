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
	uuid "github.com/satori/go.uuid"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestUpdateGameEndpoint(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "UpdateGame Endpoint EndToEnd Tests")
}

var _ = Describe("UpdateGameE2ETest Suite", func() {
	var (
		ctx     context.Context
		id      uuid.UUID
		request *dtos.UpdateGameRequestDto
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

	// "Scenario" step for testing the update game API with valid input
	Describe("Update game with valid input returns NoContent status", func() {
		BeforeEach(func() {
			request = &dtos.UpdateGameRequestDto{
				Description: gofakeit.AdjectiveDescriptive(),
				Price:       gofakeit.Price(100, 1000),
				Name:        gofakeit.Name(),
			}
		})

		// "When" step
		When("A valid request is made to update a game", func() {
			// "Then" step
			It("Should return a NoContent status", func() {
				// Create an HTTPExpect instance and make the request
				expect := httpexpect.New(GinkgoT(), integrationFixture.BaseAddress)
				expect.PUT("games/{id}").
					WithPath("id", id.String()).
					WithJSON(request).
					WithContext(ctx).
					Expect().
					Status(http.StatusNoContent)
			})
		})
	})

	// "Scenario" step for testing the update game API with invalid input
	Describe("Update game returns BadRequest with invalid input", func() {
		BeforeEach(func() {
			// Get a valid game ID from your test data
			id = uuid.NewV4()
			request = &dtos.UpdateGameRequestDto{
				Description: gofakeit.AdjectiveDescriptive(),
				Price:       0,
				Name:        gofakeit.Name(),
			}
		})
		// "When" step
		When("An invalid request is made to update a game", func() {
			// "Then" step
			It("Should return a BadRequest status", func() {
				// Create an HTTPExpect instance and make the request
				expect := httpexpect.New(GinkgoT(), integrationFixture.BaseAddress)
				expect.PUT("games/{id}").
					WithPath("id", id.String()).
					WithJSON(request).
					WithContext(context.Background()).
					Expect().
					Status(http.StatusBadRequest)
			})
		})
	})
})
