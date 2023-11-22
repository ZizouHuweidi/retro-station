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

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestSearchGamesEndpoint(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "SearchGames Endpoint EndToEnd Tests")
}

var _ = Describe("Search Games Feature", func() {
	var ctx context.Context

	_ = BeforeEach(func() {
		ctx = context.Background()

		By("Seeding the required data")
		integrationFixture.InitializeTest()
	})

	_ = AfterEach(func() {
		By("Cleanup test data")
		integrationFixture.DisposeTest()
	})

	// "Scenario" step for testing the search games API
	Describe("Search games return ok status", func() {
		// "When" step
		When("A request is made to search for games", func() {
			// "Then" step
			It("Should return an OK status", func() {
				// Create an HTTPExpect instance and make the request
				expect := httpexpect.New(GinkgoT(), integrationFixture.BaseAddress)
				expect.GET("games/search").
					WithContext(ctx).
					WithQuery("search", integrationFixture.Items[0].Name).
					Expect().
					Status(http.StatusOK)
			})
		})
	})
})
