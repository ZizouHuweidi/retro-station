package catalogs

import (
	"fmt"

	"go.opentelemetry.io/otel/metric"
	api "go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/config"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/data/uow"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/configurations/catalogs/infrastructure"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/contracts"
)

// https://pmihaylov.com/shared-components-go-microservices/
var CatalogsServiceModule = fx.Module(
	"catalogsfx",
	// Shared Modules
	config.Module,
	infrastructure.Module,

	// Features Modules
	games.Module,

	// Other provides
	fx.Provide(provideCatalogsMetrics),
	fx.Provide(uow.NewCatalogsUnitOfWork),
)

// ref: https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go
func provideCatalogsMetrics(
	cfg *config.AppOptions,
	meter metric.Meter,
) (*contracts.CatalogsMetrics, error) {
	if meter == nil {
		return nil, nil
	}

	createGameGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_create_game_grpc_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of create game grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	updateGameGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_update_game_grpc_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of update game grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	deleteGameGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_delete_game_grpc_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of delete game grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	getGameByIdGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_game_by_id_grpc_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of get game by id grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	searchGameGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_search_game_grpc_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of search game grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	createGameRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_create_game_rabbitmq_messages_total", cfg.ServiceName),
		api.WithDescription("The total number of create game rabbirmq messages"),
	)
	if err != nil {
		return nil, err
	}

	updateGameRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_update_game_rabbitmq_messages_total", cfg.ServiceName),
		api.WithDescription("The total number of update game rabbirmq messages"),
	)
	if err != nil {
		return nil, err
	}

	deleteGameRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_delete_game_rabbitmq_messages_total", cfg.ServiceName),
		api.WithDescription("The total number of delete game rabbirmq messages"),
	)
	if err != nil {
		return nil, err
	}

	successRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_search_game_rabbitmq_messages_total", cfg.ServiceName),
		api.WithDescription("The total number of success rabbitmq processed messages"),
	)
	if err != nil {
		return nil, err
	}

	errorRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_error_rabbitmq_processed_messages_total", cfg.ServiceName),
		api.WithDescription("The total number of error rabbitmq processed messages"),
	)
	if err != nil {
		return nil, err
	}

	createGameHttpRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_create_game_http_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of create game http requests"),
	)
	if err != nil {
		return nil, err
	}

	updateGameHttpRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_update_game_http_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of update game http requests"),
	)
	if err != nil {
		return nil, err
	}

	deleteGameHttpRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_delete_game_http_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of delete game http requests"),
	)
	if err != nil {
		return nil, err
	}

	getGameByIdHttpRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_game_by_id_http_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of get game by id http requests"),
	)
	if err != nil {
		return nil, err
	}

	getGamesHttpRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_games_http_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of get games http requests"),
	)
	if err != nil {
		return nil, err
	}

	searchGameHttpRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_search_game_http_requests_total", cfg.ServiceName),
		api.WithDescription("The total number of search game http requests"),
	)
	if err != nil {
		return nil, err
	}

	return &contracts.CatalogsMetrics{
		CreateGameRabbitMQMessages: createGameRabbitMQMessages,
		GetGameByIdGrpcRequests:    getGameByIdGrpcRequests,
		CreateGameGrpcRequests:     createGameGrpcRequests,
		CreateGameHttpRequests:     createGameHttpRequests,
		DeleteGameRabbitMQMessages: deleteGameRabbitMQMessages,
		DeleteGameGrpcRequests:     deleteGameGrpcRequests,
		DeleteGameHttpRequests:     deleteGameHttpRequests,
		ErrorRabbitMQMessages:      errorRabbitMQMessages,
		GetGameByIdHttpRequests:    getGameByIdHttpRequests,
		GetGamesHttpRequests:       getGamesHttpRequests,
		SearchGameGrpcRequests:     searchGameGrpcRequests,
		SearchGameHttpRequests:     searchGameHttpRequests,
		SuccessRabbitMQMessages:    successRabbitMQMessages,
		UpdateGameRabbitMQMessages: updateGameRabbitMQMessages,
		UpdateGameGrpcRequests:     updateGameGrpcRequests,
		UpdateGameHttpRequests:     updateGameHttpRequests,
	}, nil
}
