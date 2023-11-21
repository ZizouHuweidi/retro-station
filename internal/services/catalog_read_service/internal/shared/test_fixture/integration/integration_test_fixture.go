package integration

import (
	"context"
	"testing"
	"time"

	"emperror.dev/errors"
	"github.com/brianvoe/gofakeit/v6"
	rabbithole "github.com/michaelklishin/rabbit-hole"
	uuid "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/bus"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mongodb"
	config2 "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel/trace"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/config"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/app/test"
)

type IntegrationTestSharedFixture struct {
	Cfg                 *config.Config
	Log                 logger.Logger
	Bus                 bus.Bus
	GameRepository      data.GameRepository
	GameCacheRepository data.GameCacheRepository
	Container           contracts.Container
	RabbitmqCleaner     *rabbithole.Client
	rabbitmqOptions     *config2.RabbitmqOptions
	MongoOptions        *mongodb.MongoDbOptions
	BaseAddress         string
	mongoClient         *mongo.Client
	Items               []*models.Game
	Tracer              trace.Tracer
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
		result.Logger.Error(errors.WrapIf(err, "error in creating rabbithole client"))
	}

	shared := &IntegrationTestSharedFixture{
		Log:                 result.Logger,
		Container:           result.Container,
		Cfg:                 result.Cfg,
		RabbitmqCleaner:     rmqc,
		GameRepository:      result.GameRepository,
		GameCacheRepository: result.GameCacheRepository,
		Bus:                 result.Bus,
		rabbitmqOptions:     result.RabbitmqOptions,
		MongoOptions:        result.MongoDbOptions,
		BaseAddress:         result.EchoHttpOptions.BasePathAddress(),
		mongoClient:         result.MongoClient,
		Tracer:              result.Tracer,
	}

	return shared
}

func (i *IntegrationTestSharedFixture) InitializeTest() {
	i.Log.Info("InitializeTest started")

	// seed data in each test
	res, err := seedData(i.mongoClient, i.MongoOptions.Database)
	if err != nil {
		i.Log.Error(errors.WrapIf(err, "error in seeding mongodb data"))
	}

	i.Items = res
}

func (i *IntegrationTestSharedFixture) DisposeTest() {
	i.Log.Info("DisposeTest started")

	// cleanup test containers with their hooks
	if err := i.cleanupRabbitmqData(); err != nil {
		i.Log.Error(errors.WrapIf(err, "error in cleanup rabbitmq data"))
	}

	if err := i.cleanupMongoData(); err != nil {
		i.Log.Error(errors.WrapIf(err, "error in cleanup mongodb data"))
	}
}

func seedData(
	db *mongo.Client,
	databaseName string,
) ([]*models.Game, error) {
	ctx := context.Background()

	games := []*models.Game{
		{
			Id:          uuid.NewV4().String(),
			GameId:      uuid.NewV4().String(),
			Name:        gofakeit.Name(),
			CreatedAt:   time.Now(),
			Description: gofakeit.AdjectiveDescriptive(),
			Price:       gofakeit.Price(100, 1000),
			Genre:       gofakeit.RandomString([]string{"Action", "Adventure", "RPG", "Strategy", "Simulation", "Sports", "Puzzle", "Idle"}),
		},
		{
			Id:          uuid.NewV4().String(),
			GameId:      uuid.NewV4().String(),
			Name:        gofakeit.Name(),
			CreatedAt:   time.Now(),
			Description: gofakeit.AdjectiveDescriptive(),
			Price:       gofakeit.Price(100, 1000),
			Genre:       gofakeit.RandomString([]string{"Action", "Adventure", "RPG", "Strategy", "Simulation", "Sports", "Puzzle", "Idle"}),
		},
	}

	//// https://go.dev/doc/faq#convert_slice_of_interface
	gamesData := make([]interface{}, len(games))

	for i, v := range games {
		gamesData[i] = v
	}

	collection := db.Database(databaseName).Collection("games")
	_, err := collection.InsertMany(
		context.Background(),
		gamesData,
		&options.InsertManyOptions{},
	)
	if err != nil {
		return nil, errors.WrapIf(err, "error in seed database")
	}

	result, err := mongodb.Paginate[*models.Game](
		ctx,
		utils.NewListQuery(10, 1),
		collection,
		nil,
	)

	return result.Items, nil
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

func (i *IntegrationTestSharedFixture) cleanupMongoData() error {
	collections := []string{"games"}
	err := cleanupCollections(
		i.mongoClient,
		collections,
		i.MongoOptions.Database,
	)

	return err
}

func cleanupCollections(
	db *mongo.Client,
	collections []string,
	databaseName string,
) error {
	database := db.Database(databaseName)
	ctx := context.Background()

	// Iterate over the collections and delete all collections
	for _, collection := range collections {
		collection := database.Collection(collection)

		err := collection.Drop(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
