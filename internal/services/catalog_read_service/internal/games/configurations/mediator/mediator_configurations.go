package mediator

import (
	"emperror.dev/errors"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/data"
	createGameCommandV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/creating_game/v1/commands"
	createGameDtosV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/creating_game/v1/dtos"
	deleteGameCommandV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/deleting_games/v1/commands"
	getGameByIdDtosV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/get_game_by_id/v1/dtos"
	getGameByIdQueryV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/get_game_by_id/v1/queries"
	getGamesDtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/getting_games/v1/dtos"
	getGamesQueryV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/getting_games/v1/queries"
	searchGamesDtosV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/searching_games/v1/dtos"
	searchGamesQueryV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/searching_games/v1/queries"
	updateGameCommandV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/updating_games/v1/commands"
)

func ConfigGamesMediator(
	logger logger.Logger,
	mongoGameRepository data.GameRepository,
	cacheGameRepository data.GameCacheRepository,
	tracer tracing.AppTracer,
) error {
	err := mediatr.RegisterRequestHandler[*createGameCommandV1.CreateGame, *createGameDtosV1.CreateGameResponseDto](
		createGameCommandV1.NewCreateGameHandler(
			logger,
			mongoGameRepository,
			cacheGameRepository,
			tracer,
		),
	)
	if err != nil {
		return errors.WrapIf(err, "error while registering handlers in the mediator")
	}

	err = mediatr.RegisterRequestHandler[*deleteGameCommandV1.DeleteGame, *mediatr.Unit](
		deleteGameCommandV1.NewDeleteGameHandler(
			logger,
			mongoGameRepository,
			cacheGameRepository,
			tracer,
		),
	)
	if err != nil {
		return errors.WrapIf(err, "error while registering handlers in the mediator")
	}

	err = mediatr.RegisterRequestHandler[*updateGameCommandV1.UpdateGame, *mediatr.Unit](
		updateGameCommandV1.NewUpdateGameHandler(
			logger,
			mongoGameRepository,
			cacheGameRepository,
			tracer,
		),
	)
	if err != nil {
		return errors.WrapIf(err, "error while registering handlers in the mediator")
	}

	err = mediatr.RegisterRequestHandler[*getGamesQueryV1.GetGames, *getGamesDtoV1.GetGamesResponseDto](
		getGamesQueryV1.NewGetGamesHandler(logger, mongoGameRepository, tracer),
	)
	if err != nil {
		return errors.WrapIf(err, "error while registering handlers in the mediator")
	}

	err = mediatr.RegisterRequestHandler[*searchGamesQueryV1.SearchGames, *searchGamesDtosV1.SearchGamesResponseDto](
		searchGamesQueryV1.NewSearchGamesHandler(
			logger,
			mongoGameRepository,
			tracer,
		),
	)
	if err != nil {
		return errors.WrapIf(err, "error while registering handlers in the mediator")
	}

	err = mediatr.RegisterRequestHandler[*getGameByIdQueryV1.GetGameById, *getGameByIdDtosV1.GetGameByIdResponseDto](
		getGameByIdQueryV1.NewGetGameByIdHandler(
			logger,
			mongoGameRepository,
			cacheGameRepository,
			tracer,
		),
	)
	if err != nil {
		return errors.WrapIf(err, "error while registering handlers in the mediator")
	}

	return nil
}
