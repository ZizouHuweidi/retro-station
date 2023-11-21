package test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	gormPostgres "github.com/zizouhuweidi/retro-station/internal/pkg/gorm_postgres"
	"github.com/zizouhuweidi/retro-station/internal/pkg/grpc"
	config3 "github.com/zizouhuweidi/retro-station/internal/pkg/http/custom_echo/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	contracts2 "github.com/zizouhuweidi/retro-station/internal/pkg/migration/contracts"
	"github.com/zizouhuweidi/retro-station/internal/pkg/migration/goose"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/bus"
	config2 "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/containers/testcontainer/gorm"
	"github.com/zizouhuweidi/retro-station/internal/pkg/test/containers/testcontainer/rabbitmq"
	gorm2 "gorm.io/gorm"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/config"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/configurations/catalogs"
	gamesService "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc/genproto"
)

type TestApp struct{}

type TestAppResult struct {
	Cfg                     *config.AppOptions
	Bus                     bus.RabbitmqBus
	Container               contracts.Container
	Logger                  logger.Logger
	RabbitmqOptions         *config2.RabbitmqOptions
	EchoHttpOptions         *config3.EchoHttpOptions
	GormOptions             *gormPostgres.GormOptions
	CatalogUnitOfWorks      data.CatalogUnitOfWork
	GameRepository          data.GameRepository
	Gorm                    *gorm2.DB
	GameServiceClient       gamesService.GamesServiceClient
	GrpcClient              grpc.GrpcClient
	PostgresMigrationRunner contracts2.PostgresMigrationRunner
}

func NewTestApp() *TestApp {
	return &TestApp{}
}

func (a *TestApp) Run(t *testing.T) (result *TestAppResult) {
	lifetimeCtx := context.Background()

	// ref: https://github.com/uber-go/fx/blob/master/app_test.go
	appBuilder := NewCatalogsWriteTestApplicationBuilder(t)
	appBuilder.ProvideModule(catalogs.CatalogsServiceModule)
	appBuilder.ProvideModule(goose.Module)

	appBuilder.Decorate(rabbitmq.RabbitmqContainerOptionsDecorator(t, lifetimeCtx))
	appBuilder.Decorate(gorm.GormContainerOptionsDecorator(t, lifetimeCtx))

	testApp := appBuilder.Build()

	testApp.ConfigureCatalogs()

	testApp.MapCatalogsEndpoints()

	testApp.ResolveFunc(
		func(cfg *config.AppOptions,
			bus bus.RabbitmqBus,
			logger logger.Logger,
			rabbitmqOptions *config2.RabbitmqOptions,
			gormOptions *gormPostgres.GormOptions,
			catalogUnitOfWorks data.CatalogUnitOfWork,
			gameRepository data.GameRepository,
			gorm *gorm2.DB,
			echoOptions *config3.EchoHttpOptions,
			grpcClient grpc.GrpcClient,
			postgresMigrationRunner contracts2.PostgresMigrationRunner,
		) {
			grpcConnection := grpcClient.GetGrpcConnection()

			result = &TestAppResult{
				Bus:                     bus,
				Cfg:                     cfg,
				Container:               testApp,
				Logger:                  logger,
				RabbitmqOptions:         rabbitmqOptions,
				GormOptions:             gormOptions,
				GameRepository:          gameRepository,
				CatalogUnitOfWorks:      catalogUnitOfWorks,
				Gorm:                    gorm,
				EchoHttpOptions:         echoOptions,
				PostgresMigrationRunner: postgresMigrationRunner,
				GameServiceClient: gamesService.NewGamesServiceClient(
					grpcConnection,
				),
				GrpcClient: grpcClient,
			}
		},
	)
	// we need a longer timout for up and running our testcontainers
	duration := time.Second * 300

	// short timeout for handling start hooks and setup dependencies
	startCtx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	err := testApp.Start(startCtx)
	if err != nil {
		t.Errorf("Error starting, err: %v", err)
		os.Exit(1)
	}

	// waiting for grpc endpoint becomes ready in the given timeout
	err = result.GrpcClient.WaitForAvailableConnection()
	require.NoError(t, err)

	t.Cleanup(func() {
		// short timeout for handling stop hooks
		stopCtx, cancel := context.WithTimeout(context.Background(), duration)
		defer cancel()

		err = testApp.Stop(stopCtx)
		require.NoError(t, err)
	})

	return
}
