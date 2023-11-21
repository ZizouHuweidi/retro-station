package mediatr

import (
	"github.com/mehdihadeli/go-mediatr"
	logger2 "github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/producer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	createGameCommandV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/commands"
	createGameV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/dtos"
	deleteGameCommandV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/deleting_game/v1/commands"
	getGameByIdDtosV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_game_by_id/v1/dtos"
	getGameByIdQueryV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_game_by_id/v1/queries"
	getGamesDtosV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_games/v1/dtos"
	getGamesQueryV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_games/v1/queries"
	searchGamesDtosV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/searching_game/v1/dtos"
	searchGamesQueryV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/searching_game/v1/queries"
	updateGameCommandV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/commands"
)

func ConfigGamesMediator(
	logger logger2.Logger,
	uow data.CatalogUnitOfWork,
	gameRepository data.GameRepository,
	producer producer.Producer,
	tracer tracing.AppTracer,
) error {
	// https://stackoverflow.com/questions/72034479/how-to-implement-generic-interfaces
	err := mediatr.RegisterRequestHandler[*createGameCommandV1.CreateGame, *createGameV1.CreateGameResponseDto](
		createGameCommandV1.NewCreateGameHandler(logger, uow, producer, tracer),
	)
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*getGamesQueryV1.GetGames, *getGamesDtosV1.GetGamesResponseDto](
		getGamesQueryV1.NewGetGamesHandler(logger, gameRepository, tracer),
	)
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*searchGamesQueryV1.SearchGames, *searchGamesDtosV1.SearchGamesResponseDto](
		searchGamesQueryV1.NewSearchGamesHandler(logger, gameRepository, tracer),
	)
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*updateGameCommandV1.UpdateGame, *mediatr.Unit](
		updateGameCommandV1.NewUpdateGameHandler(logger, uow, producer, tracer),
	)
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*deleteGameCommandV1.DeleteGame, *mediatr.Unit](
		deleteGameCommandV1.NewDeleteGameHandler(logger, uow, producer, tracer),
	)
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*getGameByIdQueryV1.GetGameById, *getGameByIdDtosV1.GetGameByIdResponseDto](
		getGameByIdQueryV1.NewGetGameByIdHandler(logger, gameRepository, tracer),
	)
	if err != nil {
		return err
	}

	return nil
}
