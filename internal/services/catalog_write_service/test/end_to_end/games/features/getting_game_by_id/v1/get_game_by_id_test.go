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

func TestGetGameByIdEndpoint(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "GetGameById Endpoint EndToEnd Tests")
}

var _ = Describe("Get Game By Id Feature", func() {
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

	// "Scenario" step for testing the get game by ID API with a valid ID
	Describe("Get game by ID with a valid ID returns ok status", func() {
		// "When" step
		When("A valid request is made with a valid ID", func() {
			// "Then" step
			It("Should return an OK status", func() {
				expect := httpexpect.New(GinkgoT(), integrationFixture.BaseAddress)
				expect.GET("games/{id}").
					WithPath("id", id).
					WithContext(ctx).
					Expect().
					Status(http.StatusOK)
			})
		})
	})

	// "Scenario" step for testing the get game by ID API with a valid ID
	Describe("Get game by ID with a invalid ID returns NotFound status", func() {
		BeforeEach(func() {
			// Generate an invalid UUID
			id = uuid.NewV4()
		})
		When("An invalid request is made with an invalid ID", func() {
			// "Then" step
			It("Should return a NotFound status", func() {
				// Create an HTTPExpect instance and make the request
				expect := httpexpect.New(GinkgoT(), integrationFixture.BaseAddress)
				expect.GET("games/{id}").
					WithPath("id", id.String()).
					WithContext(ctx).
					Expect().
					Status(http.StatusNotFound)
			})
		})
	})
})
