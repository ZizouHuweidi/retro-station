package unit_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	defaultLogger "github.com/zizouhuweidi/retro-station/internal/pkg/logger/default_logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	mocks3 "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/mocks"
	"go.opentelemetry.io/otel/trace"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/config"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	dto "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/mocks/testData"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/mocks"
)

type UnitTestSharedFixture struct {
	Cfg *config.AppOptions
	Log logger.Logger
	suite.Suite
	Items          []*models.Game
	Uow            *mocks.CatalogUnitOfWork
	GameRepository *mocks.GameRepository
	Bus            *mocks3.Bus
	Tracer         trace.Tracer
}

func NewUnitTestSharedFixture(t *testing.T) *UnitTestSharedFixture {
	// we could use EmptyLogger if we don't want to log anything
	defaultLogger.SetupDefaultLogger()
	log := defaultLogger.Logger
	cfg := &config.AppOptions{}

	err := configMapper()
	require.NoError(t, err)

	// empty tracer, just for testing
	nopetracer := trace.NewNoopTracerProvider()
	testTracer := nopetracer.Tracer("test_tracer")

	unit := &UnitTestSharedFixture{
		Cfg:    cfg,
		Log:    log,
		Items:  testData.Games,
		Tracer: testTracer,
	}

	return unit
}

func configMapper() error {
	err := mapper.CreateMap[*models.Game, *dto.GameDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*dto.GameDto, *models.Game]()
	if err != nil {
		return err
	}

	return nil
}

// //////////////Shared Hooks////////////////
func (c *UnitTestSharedFixture) SetupTest() {
	// create new mocks
	gameRepository := &mocks.GameRepository{}
	bus := &mocks3.Bus{}
	uow := &mocks.CatalogUnitOfWork{}
	catalogContext := &mocks.CatalogContext{}

	//// or just clear the mocks
	//c.Bus.ExpectedCalls = nil
	//c.Bus.Calls = nil
	//c.Uow.ExpectedCalls = nil
	//c.Uow.Calls = nil
	//c.GameRepository.ExpectedCalls = nil
	//c.GameRepository.Calls = nil

	uow.On("Games").Return(gameRepository)
	catalogContext.On("Games").Return(gameRepository)

	var mockUOW *mock.Call
	mockUOW = uow.On("Do", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			fn, ok := args.Get(1).(data.CatalogUnitOfWorkActionFunc)
			if !ok {
				panic("argument mismatch")
			}
			fmt.Println(fn)

			mockUOW.Return(fn(catalogContext))
		})

	mockUOW.Times(1)
	bus.On("PublishMessage", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	c.Uow = uow
	c.GameRepository = gameRepository
	c.Bus = bus
}

func (c *UnitTestSharedFixture) CleanupMocks() {
	c.SetupTest()
}

func (c *UnitTestSharedFixture) TearDownSuite() {
	mapper.ClearMappings()
}
