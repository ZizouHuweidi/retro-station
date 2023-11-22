//go:build integration
// +build integration

package uow

import (
	"context"
	"testing"
	"time"

	"emperror.dev/errors"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	data2 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

var integrationFixture *integration.IntegrationTestSharedFixture

func TestUnitOfWork(t *testing.T) {
	RegisterFailHandler(Fail)
	integrationFixture = integration.NewIntegrationTestSharedFixture(t)
	RunSpecs(t, "CatalogsUnitOfWork Integration Tests")
}

var _ = Describe("CatalogsUnitOfWork Feature", func() {
	// Define variables to hold repository and game data
	var (
		ctx   context.Context
		err   error
		games *utils.ListResult[*models.Game]
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

	// "Scenario" step for testing a UnitOfWork action that should roll back on error
	Describe("Rollback on error", func() {
		// "When" step
		When("The UnitOfWork Do executed and there is an error in the execution", func() {
			It("Should roll back the changes and not affect the database", func() {
				err = integrationFixture.CatalogUnitOfWorks.Do(ctx, func(catalogContext data2.CatalogContext) error {
					_, err := catalogContext.Games().CreateGame(ctx,
						&models.Game{
							Name:        gofakeit.Name(),
							Description: gofakeit.AdjectiveDescriptive(),
							GameId:      uuid.NewV4(),
							Price:       gofakeit.Price(100, 1000),
							CreatedAt:   time.Now(),
						})
					Expect(err).NotTo(HaveOccurred()) // Successful game creation

					return errors.New("error rollback")
				})
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring("error rollback")))

				games, err := integrationFixture.GameRepository.GetAllGames(ctx, utils.NewListQuery(10, 1))
				Expect(err).To(BeNil())

				Expect(len(games.Items)).To(Equal(2)) // Ensure no changes in the database
			})
		})
	})

	// "Scenario" step for testing a UnitOfWork action that should rollback on panic
	Describe("Rollback on panic", func() {
		// "When" step
		When("The UnitOfWork Do executed and there is an panic in the execution", func() {
			It("Should roll back the changes and not affect the database", func() {
				err = integrationFixture.CatalogUnitOfWorks.Do(ctx, func(catalogContext data2.CatalogContext) error {
					_, err := catalogContext.Games().CreateGame(ctx,
						&models.Game{
							Name:        gofakeit.Name(),
							Description: gofakeit.AdjectiveDescriptive(),
							GameId:      uuid.NewV4(),
							Price:       gofakeit.Price(100, 1000),
							CreatedAt:   time.Now(),
						})
					Expect(err).To(BeNil()) // Successful game creation

					panic(errors.New("panic rollback"))
				})
				Expect(err).To(HaveOccurred())

				games, err = integrationFixture.GameRepository.GetAllGames(ctx, utils.NewListQuery(10, 1))
				Expect(err).To(BeNil())

				Expect(len(games.Items)).To(Equal(2)) // Ensure no changes in the database
			})
		})
	})

	// "Scenario" step for testing a UnitOfWork action that should rollback when the context is canceled
	Describe("Cancelling the context", func() {
		// "When" step
		When("the UnitOfWork Do executed and cancel the context", func() {
			It("Should roll back the changes and not affect the database", func() {
				cancelCtx, cancel := context.WithCancel(ctx)

				err := integrationFixture.CatalogUnitOfWorks.Do(
					cancelCtx,
					func(catalogContext data2.CatalogContext) error {
						_, err := catalogContext.Games().CreateGame(ctx,
							&models.Game{
								Name:        gofakeit.Name(),
								Description: gofakeit.AdjectiveDescriptive(),
								GameId:      uuid.NewV4(),
								Price:       gofakeit.Price(100, 1000),
								CreatedAt:   time.Now(),
							})
						Expect(err).To(BeNil()) // Successful game creation

						_, err = catalogContext.Games().CreateGame(ctx,
							&models.Game{
								Name:        gofakeit.Name(),
								Description: gofakeit.AdjectiveDescriptive(),
								GameId:      uuid.NewV4(),
								Price:       gofakeit.Price(100, 1000),
								CreatedAt:   time.Now(),
							})
						Expect(err).To(BeNil()) // Successful game creation

						cancel() // Cancel the context

						return err
					},
				)
				Expect(err).To(HaveOccurred())

				// Validate that changes are rolled back in the database
				games, err := integrationFixture.GameRepository.GetAllGames(ctx, utils.NewListQuery(10, 1))
				Expect(err).To(BeNil())
				Expect(len(games.Items)).To(Equal(2)) // Ensure no changes in the database
			})
		})
	})

	// "Scenario" step for testing a UnitOfWork action that should commit on success
	Describe("Commit on success", func() {
		// "When" step
		When("the UnitOfWork Do executed and operation was successfull", func() {
			It("Should commit the changes to the database", func() {
				err := integrationFixture.CatalogUnitOfWorks.Do(ctx, func(catalogContext data2.CatalogContext) error {
					_, err := catalogContext.Games().CreateGame(ctx,
						&models.Game{
							Name:        gofakeit.Name(),
							Description: gofakeit.AdjectiveDescriptive(),
							GameId:      uuid.NewV4(),
							Price:       gofakeit.Price(100, 1000),
							CreatedAt:   time.Now(),
						})
					Expect(err).To(BeNil()) // Successful game creation

					return err
				})
				Expect(err).To(BeNil()) // No error indicates success

				// Validate that changes are committed in the database
				games, err := integrationFixture.GameRepository.GetAllGames(ctx, utils.NewListQuery(10, 1))
				Expect(err).To(BeNil())
				Expect(len(games.Items)).To(Equal(3)) // Ensure changes in the database
			})
		})
	})
})
