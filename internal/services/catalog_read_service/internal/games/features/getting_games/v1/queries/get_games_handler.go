package queries

import (
	"context"

	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/dto"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/getting_games/v1/dtos"
)

type GetGamesHandler struct {
	log             logger.Logger
	mongoRepository data.GameRepository
	tracer          tracing.AppTracer
}

func NewGetGamesHandler(
	log logger.Logger,
	mongoRepository data.GameRepository,
	tracer tracing.AppTracer,
) *GetGamesHandler {
	return &GetGamesHandler{log: log, mongoRepository: mongoRepository, tracer: tracer}
}

func (c *GetGamesHandler) Handle(
	ctx context.Context,
	query *GetGames,
) (*dtos.GetGamesResponseDto, error) {
	ctx, span := c.tracer.Start(ctx, "GetGamesHandler.Handle")
	span.SetAttributes(attribute.Object("Query", query))
	defer span.End()

	games, err := c.mongoRepository.GetAllGames(ctx, query.ListQuery)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[GetGamesHandler_Handle.GetAllGames] error in getting games in the repository",
			),
		)
	}

	listResultDto, err := utils.ListResultToListResultDto[*dto.GameDto](games)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[GetGamesHandler_Handle.ListResultToListResultDto] error in the mapping ListResultToListResultDto",
			),
		)
	}

	c.log.Info("[GetGamesHandler.Handle] games fetched")

	return &dtos.GetGamesResponseDto{Games: listResultDto}, nil
}
