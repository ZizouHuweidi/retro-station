package contracts

import (
	"go.opentelemetry.io/otel/metric"
)

type CatalogsMetrics struct {
	CreateGameGrpcRequests     metric.Float64Counter
	UpdateGameGrpcRequests     metric.Float64Counter
	DeleteGameGrpcRequests     metric.Float64Counter
	GetGameByIdGrpcRequests    metric.Float64Counter
	SearchGameGrpcRequests     metric.Float64Counter
	CreateGameHttpRequests     metric.Float64Counter
	UpdateGameHttpRequests     metric.Float64Counter
	DeleteGameHttpRequests     metric.Float64Counter
	GetGameByIdHttpRequests    metric.Float64Counter
	GetGamesHttpRequests       metric.Float64Counter
	SearchGameHttpRequests     metric.Float64Counter
	SuccessRabbitMQMessages    metric.Float64Counter
	ErrorRabbitMQMessages      metric.Float64Counter
	CreateGameRabbitMQMessages metric.Float64Counter
	UpdateGameRabbitMQMessages metric.Float64Counter
	DeleteGameRabbitMQMessages metric.Float64Counter
}
