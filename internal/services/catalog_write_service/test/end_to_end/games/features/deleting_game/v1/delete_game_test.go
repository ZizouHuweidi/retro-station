//go:build e2e
// +build e2e

package v1

import (
	"context"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestDeleteGameEndpoint(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "DeleteGame Endpoint EndToEnd Tests")
}

var _ = Describe("Delete Game Feature", func() {
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

	// "Scenario" step for testing the delete game API with valid input
	Describe("Delete game with valid input returns NoContent status", func() {
		// "When" step
		When("A valid request is made to delete a game", func() {
			// "Then" step
			It("Should return a NoContent status", func() {
				// Create an HTTPExpect instance and make the request
				expect := httpexpect.New(GinkgoT(), integrationFixture.BaseAddress)
				expect.DELETE("games/{id}").
					WithContext(ctx).
					WithPath("id", id.String()).
					Expect().
					Status(http.StatusNoContent)
			})
		})
	})

	// "Scenario" step for testing the delete game API with invalid ID
	Describe("Delete game with with invalid ID returns NotFound status", func() {
		BeforeEach(func() {
			// Generate an invalid UUID
			id = uuid.NewV4()
		})

		// "When" step
		When("An invalid request is made with an invalid ID", func() {
			// "Then" step
			It("Should return a NotFound status", func() {
				// Create an HTTPExpect instance and make the request
				expect := httpexpect.New(GinkgoT(), integrationFixture.BaseAddress)
				expect.DELETE("games/{id}").
					WithContext(ctx).
					WithPath("id", id.String()).
					Expect().
					Status(http.StatusNotFound)
			})
		})
	})
})
