package integration

import (
	"context"
	"testing"
	"time"

	"emperror.dev/errors"
	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/lib/pq"
	rabbithole "github.com/michaelklishin/rabbit-hole"
	uuid "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	gormPostgres "github.com/zizouhuweidi/retro-station/internal/pkg/gorm_postgres"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/bus"
	config2 "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/testfixture"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gorm.io/gorm"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/config"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/app/test"
	gamesService "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc/genproto"
)

type IntegrationTestSharedFixture struct {
	Cfg                *config.AppOptions
	Log                logger.Logger
	Bus                bus.Bus
	CatalogUnitOfWorks data.CatalogUnitOfWork
	GameRepository     data.GameRepository
	Container          contracts.Container
	DbCleaner          dbcleaner.DbCleaner
	RabbitmqCleaner    *rabbithole.Client
	rabbitmqOptions    *config2.RabbitmqOptions
	Gorm               *gorm.DB
	BaseAddress        string
	Items              []*models.Game
	GameServiceClient  gamesService.GamesServiceClient
}

func NewIntegrationTestSharedFixture(
	t *testing.T,
) *IntegrationTestSharedFixture {
	result := test.NewTestApp().Run(t)

	// https://github.com/michaelklishin/rabbit-hole
	rmqc, err := rabbithole.NewClient(
		result.RabbitmqOptions.RabbitmqHostOptions.HttpEndPoint(),
		result.RabbitmqOptions.RabbitmqHostOptions.UserName,
		result.RabbitmqOptions.RabbitmqHostOptions.Password)
	if err != nil {
		result.Logger.Error(
			errors.WrapIf(err, "error in creating rabbithole client"),
		)
	}

	shared := &IntegrationTestSharedFixture{
		Log:                result.Logger,
		Container:          result.Container,
		Cfg:                result.Cfg,
		RabbitmqCleaner:    rmqc,
		GameRepository:     result.GameRepository,
		CatalogUnitOfWorks: result.CatalogUnitOfWorks,
		Bus:                result.Bus,
		rabbitmqOptions:    result.RabbitmqOptions,
		Gorm:               result.Gorm,
		BaseAddress:        result.EchoHttpOptions.BasePathAddress(),
		GameServiceClient:  result.GameServiceClient,
	}

	migrateDatabase(result)

	return shared
}

func (i *IntegrationTestSharedFixture) InitializeTest() {
	i.Log.Info("InitializeTest started")

	// seed data in each test
	res, err := seedData(i.Gorm)
	if err != nil {
		i.Log.Error(errors.WrapIf(err, "error in seeding data in postgres"))
	}

	i.Items = res
}

func (i *IntegrationTestSharedFixture) DisposeTest() {
	i.Log.Info("DisposeTest started")

	// cleanup test containers with their hooks
	if err := i.cleanupRabbitmqData(); err != nil {
		i.Log.Error(errors.WrapIf(err, "error in cleanup rabbitmq data"))
	}

	if err := i.cleanupPostgresData(); err != nil {
		i.Log.Error(errors.WrapIf(err, "error in cleanup postgres data"))
	}
}

func (i *IntegrationTestSharedFixture) cleanupRabbitmqData() error {
	// https://github.com/michaelklishin/rabbit-hole
	// Get all queues
	queues, err := i.RabbitmqCleaner.ListQueuesIn(
		i.rabbitmqOptions.RabbitmqHostOptions.VirtualHost,
	)
	if err != nil {
		return err
	}

	// clear each queue
	for _, queue := range queues {
		_, err = i.RabbitmqCleaner.PurgeQueue(
			i.rabbitmqOptions.RabbitmqHostOptions.VirtualHost,
			queue.Name,
		)

		return err
	}

	return nil
}

func (i *IntegrationTestSharedFixture) cleanupPostgresData() error {
	tables := []string{"games"}
	// Iterate over the tables and delete all records
	for _, table := range tables {
		err := i.Gorm.Exec("DELETE FROM " + table).Error

		return err
	}

	return nil
}

func seedData(gormDB *gorm.DB) ([]*models.Game, error) {
	games := []*models.Game{
		{
			GameId:      uuid.NewV4(),
			Name:        gofakeit.Name(),
			CreatedAt:   time.Now(),
			Description: gofakeit.AdjectiveDescriptive(),
			Price:       gofakeit.Price(100, 1000),
		},
		{
			GameId:      uuid.NewV4(),
			Name:        gofakeit.Name(),
			CreatedAt:   time.Now(),
			Description: gofakeit.AdjectiveDescriptive(),
			Price:       gofakeit.Price(100, 1000),
		},
	}

	// migration will do in app configuration
	// seed data
	err := gormDB.CreateInBatches(games, len(games)).Error
	if err != nil {
		return nil, errors.Wrap(err, "error in seed database")
	}

	return games, nil
}

func seedAndMigration(gormDB *gorm.DB) ([]*models.Game, error) {
	// migration
	err := gormDB.AutoMigrate(models.Game{})
	if err != nil {
		return nil, errors.WrapIf(err, "error in seed database")
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, errors.WrapIf(err, "error in seed database")
	}

	// https://github.com/go-testfixtures/testfixtures#templating
	// seed data
	var data []struct {
		Name        string
		GameId      uuid.UUID
		Description string
	}

	f := []struct {
		Name        string
		GameId      uuid.UUID
		Description string
	}{
		{gofakeit.Name(), uuid.NewV4(), gofakeit.AdjectiveDescriptive()},
		{gofakeit.Name(), uuid.NewV4(), gofakeit.AdjectiveDescriptive()},
	}

	data = append(data, f...)

	err = testfixture.RunPostgresFixture(
		db,
		[]string{"db/fixtures/games"},
		map[string]interface{}{
			"Games": data,
		})
	if err != nil {
		return nil, errors.WrapIf(err, "error in seed database")
	}

	result, err := gormPostgres.Paginate[*models.Game](
		context.Background(),
		utils.NewListQuery(10, 1),
		gormDB,
	)
	return result.Items, nil
}

func migrateDatabase(result *test.TestAppResult) {
	err := result.PostgresMigrationRunner.Up(context.Background(), 0)
	if err != nil {
		result.Logger.Fatalf("error in catalog_service migration, err: %s", err)
	}
}
