package getGameByIdQuery

import (
	"context"
	"fmt"
	"net/http"

	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_game_by_id/v1/dtos"
)

type GetGameByIdHandler struct {
	log    logger.Logger
	pgRepo data.GameRepository
	tracer tracing.AppTracer
}

func NewGetGameByIdHandler(
	log logger.Logger,
	pgRepo data.GameRepository,
	tracer tracing.AppTracer,
) *GetGameByIdHandler {
	return &GetGameByIdHandler{log: log, pgRepo: pgRepo, tracer: tracer}
}

func (q *GetGameByIdHandler) Handle(
	ctx context.Context,
	query *GetGameById,
) (*dtos.GetGameByIdResponseDto, error) {
	ctx, span := q.tracer.Start(ctx, "GetGameByIdHandler.Handle")
	span.SetAttributes(attribute.Object("Query", query))
	span.SetAttributes(attribute2.String("GameId", query.GameID.String()))
	defer span.End()

	game, err := q.pgRepo.GetGameById(ctx, query.GameID)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrapWithCode(
				err,
				http.StatusNotFound,
				fmt.Sprintf(
					"[GetGameByIdHandler_Handle.GetGameById] error in getting game with id %s in the repository",
					query.GameID.String(),
				),
			),
		)
	}

	gameDto, err := mapper.Map[*dtoV1.GameDto](game)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[GetGameByIdHandler_Handle.Map] error in the mapping game",
			),
		)
	}

	q.log.Infow(
		fmt.Sprintf(
			"[GetGameByIdHandler.Handle] game with id: {%s} fetched",
			query.GameID,
		),
		logger.Fields{"GameId": query.GameID.String()},
	)

	return &dtos.GetGameByIdResponseDto{Game: gameDto}, nil
}
