//go:build integration
// +build integration

package v1

import (
	"context"
	"testing"
	"time"

	"github.com/mehdihadeli/go-mediatr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_games/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_games/v1/queries"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestGetGames(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "Get Games Integration Tests")
}

var _ = Describe("Get All Games Feature", func() {
	// Define variables to hold query and result data
	var (
		ctx         context.Context
		query       *queries.GetGames
		queryResult *dtos.GetGamesResponseDto
		err         error
	)

	_ = BeforeEach(func() {
		By("Seeding the required data")
		integrationFixture.InitializeTest()
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

	// "Scenario" step for testing getting a list of existing games
	Describe("Getting a list of existing games from the database", func() {
		Context("Given existing games in the database", func() {
			BeforeEach(func() {
				// Create a query to retrieve a list of games
				query, err = queries.NewGetGames(utils.NewListQuery(10, 1))
				Expect(err).To(BeNil())
			})

			// "When" step
			When("the GteGames query is executed for existing games", func() {
				BeforeEach(func() {
					queryResult, err = mediatr.Send[*queries.GetGames, *dtos.GetGamesResponseDto](ctx, query)
				})

				// "Then" step
				It("Should not return an error", func() {
					Expect(err).To(BeNil())
				})

				It("Should return a non-nil result", func() {
					Expect(queryResult).NotTo(BeNil())
				})

				It("Should return a list of games with items", func() {
					Expect(queryResult.Games).NotTo(BeNil())
					Expect(queryResult.Games.Items).NotTo(BeEmpty())
				})

				It("Should return the expected number of games", func() {
					// Replace 'len(c.Items)' with the expected number of games
					Expect(len(queryResult.Games.Items)).To(Equal(len(integrationFixture.Items)))
				})
			})
		})
	})
})
