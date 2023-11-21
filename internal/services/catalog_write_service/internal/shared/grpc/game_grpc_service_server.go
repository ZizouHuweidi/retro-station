package grpc

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"github.com/mehdihadeli/go-mediatr"
	uuid "github.com/satori/go.uuid"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	attribute2 "go.opentelemetry.io/otel/attribute"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	createGameCommandV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/commands"
	createGameDtosV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/dtos"
	getGameByIdDtosV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_game_by_id/v1/dtos"
	getGameByIdQueryV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_game_by_id/v1/queries"
	updateGameCommandV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/contracts"
	gamesService "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc/genproto"
)

var grpcMetricsAttr = api.WithAttributes(
	attribute2.Key("MetricsType").String("Http"),
)

type GameGrpcServiceServer struct {
	catalogsMetrics *contracts.CatalogsMetrics
	logger          logger.Logger
	// Ref:https://github.com/grpc/grpc-go/issues/3794#issuecomment-720599532
	// game_service_client.UnimplementedGamesServiceServer
}

func NewGameGrpcService(
	catalogsMetrics *contracts.CatalogsMetrics,
	logger logger.Logger,
) *GameGrpcServiceServer {
	return &GameGrpcServiceServer{catalogsMetrics: catalogsMetrics, logger: logger}
}

func (s *GameGrpcServiceServer) CreateGame(
	ctx context.Context,
	req *gamesService.CreateGameReq,
) (*gamesService.CreateGameRes, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.Object("Request", req))
	s.catalogsMetrics.CreateGameGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	command, err := createGameCommandV1.NewCreateGame(
		req.GetName(),
		req.GetDescription(),
		req.GetPrice(),
		req.GetGenre(),
	)
	if err != nil {
		validationErr := customErrors.NewValidationErrorWrap(
			err,
			"[GameGrpcServiceServer_CreateGame.StructCtx] command validation failed",
		)
		s.logger.Errorf(
			fmt.Sprintf(
				"[GameGrpcServiceServer_CreateGame.StructCtx] err: %v",
				validationErr,
			),
		)
		return nil, validationErr
	}

	result, err := mediatr.Send[*createGameCommandV1.CreateGame, *createGameDtosV1.CreateGameResponseDto](
		ctx,
		command,
	)
	if err != nil {
		err = errors.WithMessage(
			err,
			"[GameGrpcServiceServer_CreateGame.Send] error in sending CreateGame",
		)
		s.logger.Errorw(
			fmt.Sprintf(
				"[GameGrpcServiceServer_CreateGame.Send] id: {%s}, err: %v",
				command.GameID,
				err,
			),
			logger.Fields{"GameId": command.GameID},
		)
		return nil, err
	}

	return &gamesService.CreateGameRes{GameId: result.GameID.String()}, nil
}

func (s *GameGrpcServiceServer) UpdateGame(
	ctx context.Context,
	req *gamesService.UpdateGameReq,
) (*gamesService.UpdateGameRes, error) {
	s.catalogsMetrics.UpdateGameGrpcRequests.Add(ctx, 1, grpcMetricsAttr)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.Object("Request", req))

	gameUUID, err := uuid.FromString(req.GetGameId())
	if err != nil {
		badRequestErr := customErrors.NewBadRequestErrorWrap(
			err,
			"[GameGrpcServiceServer_UpdateGame.uuid.FromString] error in converting uuid",
		)
		s.logger.Errorf(
			fmt.Sprintf(
				"[GameGrpcServiceServer_UpdateGame.uuid.FromString] err: %v",
				badRequestErr,
			),
		)
		return nil, badRequestErr
	}

	command, err := updateGameCommandV1.NewUpdateGame(
		gameUUID,
		req.GetName(),
		req.GetDescription(),
		req.GetPrice(),
		req.GetGenre(),
	)
	if err != nil {
		validationErr := customErrors.NewValidationErrorWrap(
			err,
			"[GameGrpcServiceServer_UpdateGame.StructCtx] command validation failed",
		)
		s.logger.Errorf(
			fmt.Sprintf(
				"[GameGrpcServiceServer_UpdateGame.StructCtx] err: %v",
				validationErr,
			),
		)
		return nil, validationErr
	}

	if _, err = mediatr.Send[*updateGameCommandV1.UpdateGame, *mediatr.Unit](ctx, command); err != nil {
		err = errors.WithMessage(
			err,
			"[GameGrpcServiceServer_UpdateGame.Send] error in sending CreateGame",
		)
		s.logger.Errorw(
			fmt.Sprintf(
				"[GameGrpcServiceServer_UpdateGame.Send] id: {%s}, err: %v",
				command.GameID,
				err,
			),
			logger.Fields{"GameId": command.GameID},
		)
		return nil, err
	}

	return &gamesService.UpdateGameRes{}, nil
}

func (s *GameGrpcServiceServer) GetGameById(
	ctx context.Context,
	req *gamesService.GetGameByIdReq,
) (*gamesService.GetGameByIdRes, error) {
	//// we could use trace manually, but I used grpc middleware for doing this
	//ctx, span, clean := grpcTracing.StartGrpcServerTracerSpan(ctx, "GameGrpcServiceServer.GetGameById")
	//defer clean()

	s.catalogsMetrics.GetGameByIdGrpcRequests.Add(ctx, 1, grpcMetricsAttr)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.Object("Request", req))

	gameUUID, err := uuid.FromString(req.GetGameId())
	if err != nil {
		badRequestErr := customErrors.NewBadRequestErrorWrap(
			err,
			"[GameGrpcServiceServer_GetGameById.uuid.FromString] error in converting uuid",
		)
		s.logger.Errorf(
			fmt.Sprintf(
				"[GameGrpcServiceServer_GetGameById.uuid.FromString] err: %v",
				badRequestErr,
			),
		)
		return nil, badRequestErr
	}

	query, err := getGameByIdQueryV1.NewGetGameById(gameUUID)
	if err != nil {
		validationErr := customErrors.NewValidationErrorWrap(
			err,
			"[GameGrpcServiceServer_GetGameById.StructCtx] query validation failed",
		)
		s.logger.Errorf(
			fmt.Sprintf(
				"[GameGrpcServiceServer_GetGameById.StructCtx] err: %v",
				validationErr,
			),
		)
		return nil, validationErr
	}

	queryResult, err := mediatr.Send[*getGameByIdQueryV1.GetGameById, *getGameByIdDtosV1.GetGameByIdResponseDto](
		ctx,
		query,
	)
	if err != nil {
		err = errors.WithMessage(
			err,
			"[GameGrpcServiceServer_GetGameById.Send] error in sending GetGameById",
		)
		s.logger.Errorw(
			fmt.Sprintf(
				"[GameGrpcServiceServer_GetGameById.Send] id: {%s}, err: %v",
				query.GameID,
				err,
			),
			logger.Fields{"GameId": query.GameID},
		)
		return nil, err
	}

	game, err := mapper.Map[*gamesService.Game](queryResult.Game)
	if err != nil {
		err = errors.WithMessage(
			err,
			"[GameGrpcServiceServer_GetGameById.Map] error in mapping game",
		)
		return nil, err
	}

	return &gamesService.GetGameByIdRes{Game: game}, nil
}
