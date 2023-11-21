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
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/searching_games/v1/dtos"
)

type SearchGamesHandler struct {
	log             logger.Logger
	mongoRepository data.GameRepository
	tracer          tracing.AppTracer
}

func NewSearchGamesHandler(
	log logger.Logger,
	repository data.GameRepository,
	tracer tracing.AppTracer,
) *SearchGamesHandler {
	return &SearchGamesHandler{log: log, mongoRepository: repository, tracer: tracer}
}

func (c *SearchGamesHandler) Handle(
	ctx context.Context,
	query *SearchGames,
) (*dtos.SearchGamesResponseDto, error) {
	ctx, span := c.tracer.Start(ctx, "SearchGamesHandler.Handle")
	span.SetAttributes(attribute.Object("Query", query))
	defer span.End()

	games, err := c.mongoRepository.SearchGames(ctx, query.SearchText, query.ListQuery)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[SearchGamesHandler_Handle.SearchGames] error in searching games in the repository",
			),
		)
	}

	listResultDto, err := utils.ListResultToListResultDto[*dto.GameDto](games)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			customErrors.NewApplicationErrorWrap(
				err,
				"[SearchGamesHandler_Handle.ListResultToListResultDto] error in the mapping ListResultToListResultDto",
			),
		)
	}
	c.log.Info("[SearchGamesHandler.Handle] games fetched")

	return &dtos.SearchGamesResponseDto{Games: listResultDto}, nil
}
