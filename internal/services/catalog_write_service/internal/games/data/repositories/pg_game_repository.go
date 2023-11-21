package repositories

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/retro-station/internal/pkg/core/data"
	"github.com/zizouhuweidi/retro-station/internal/pkg/gorm_postgres/repository"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	attribute2 "go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"

	data2 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
)

type postgresGameRepository struct {
	log                   logger.Logger
	gormGenericRepository data.GenericRepository[*models.Game]
	tracer                tracing.AppTracer
}

func NewPostgresGameRepository(
	log logger.Logger,
	db *gorm.DB,
	tracer tracing.AppTracer,
) data2.GameRepository {
	gormRepository := repository.NewGenericGormRepository[*models.Game](db)
	return &postgresGameRepository{
		log:                   log,
		gormGenericRepository: gormRepository,
		tracer:                tracer,
	}
}

func (p *postgresGameRepository) GetAllGames(
	ctx context.Context,
	listQuery *utils.ListQuery,
) (*utils.ListResult[*models.Game], error) {
	ctx, span := p.tracer.Start(ctx, "postgresGameRepository.GetAllGames")
	defer span.End()

	result, err := p.gormGenericRepository.GetAll(ctx, listQuery)
	if err != nil {
		return nil, tracing.TraceErrFromContext(
			ctx,
			errors.WrapIf(
				err,
				"[postgresGameRepository_GetAllGames.Paginate] error in the paginate",
			),
		)
	}

	p.log.Infow(
		"[postgresGameRepository.GetAllGames] games loaded",
		logger.Fields{"GamesResult": result},
	)
	span.SetAttributes(attribute.Object("GamesResult", result))

	return result, nil
}

func (p *postgresGameRepository) SearchGames(
	ctx context.Context,
	searchText string,
	listQuery *utils.ListQuery,
) (*utils.ListResult[*models.Game], error) {
	ctx, span := p.tracer.Start(ctx, "postgresGameRepository.SearchGames")
	span.SetAttributes(attribute2.String("SearchText", searchText))
	defer span.End()

	result, err := p.gormGenericRepository.Search(ctx, searchText, listQuery)
	if err != nil {
		return nil, tracing.TraceErrFromContext(
			ctx,
			errors.WrapIf(
				err,
				"[postgresGameRepository_SearchGames.Paginate] error in the paginate",
			),
		)
	}

	p.log.Infow(
		fmt.Sprintf(
			"[postgresGameRepository.SearchGames] games loaded for search term '%s'",
			searchText,
		),
		logger.Fields{"GamesResult": result},
	)
	span.SetAttributes(attribute.Object("GamesResult", result))

	return result, nil
}

func (p *postgresGameRepository) GetGameById(
	ctx context.Context,
	uuid uuid.UUID,
) (*models.Game, error) {
	ctx, span := p.tracer.Start(ctx, "postgresGameRepository.GetGameById")
	span.SetAttributes(attribute2.String("GameId", uuid.String()))
	defer span.End()

	game, err := p.gormGenericRepository.GetById(ctx, uuid)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"[postgresGameRepository_GetGameById.First] can't find the game with id %s into the database.",
					uuid,
				),
			),
		)
	}

	span.SetAttributes(attribute.Object("Game", game))
	p.log.Infow(
		fmt.Sprintf(
			"[postgresGameRepository.GetGameById] game with id %s laoded",
			uuid.String(),
		),
		logger.Fields{"Game": game, "GameId": uuid},
	)

	return game, nil
}

func (p *postgresGameRepository) CreateGame(
	ctx context.Context,
	game *models.Game,
) (*models.Game, error) {
	ctx, span := p.tracer.Start(ctx, "postgresGameRepository.CreateGame")
	defer span.End()

	err := p.gormGenericRepository.Add(ctx, game)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				"[postgresGameRepository_CreateGame.Create] error in the inserting game into the database.",
			),
		)
	}

	span.SetAttributes(attribute.Object("Game", game))
	p.log.Infow(
		fmt.Sprintf(
			"[postgresGameRepository.CreateGame] game with id '%s' created",
			game.GameId,
		),
		logger.Fields{"Game": game, "GameId": game.GameId},
	)

	return game, nil
}

func (p *postgresGameRepository) UpdateGame(
	ctx context.Context,
	updateGame *models.Game,
) (*models.Game, error) {
	ctx, span := p.tracer.Start(ctx, "postgresGameRepository.UpdateGame")
	defer span.End()

	err := p.gormGenericRepository.Update(ctx, updateGame)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"[postgresGameRepository_UpdateGame.Save] error in updating game with id %s into the database.",
					updateGame.GameId,
				),
			),
		)
	}

	span.SetAttributes(attribute.Object("Game", updateGame))
	p.log.Infow(
		fmt.Sprintf(
			"[postgresGameRepository.UpdateGame] game with id '%s' updated",
			updateGame.GameId,
		),
		logger.Fields{"Game": updateGame, "GameId": updateGame.GameId},
	)

	return updateGame, nil
}

func (p *postgresGameRepository) DeleteGameByID(ctx context.Context, uuid uuid.UUID) error {
	ctx, span := p.tracer.Start(ctx, "postgresGameRepository.UpdateGame")
	span.SetAttributes(attribute2.String("GameId", uuid.String()))
	defer span.End()

	err := p.gormGenericRepository.Delete(ctx, uuid)
	if err != nil {
		return tracing.TraceErrFromSpan(span, errors.WrapIf(err, fmt.Sprintf(
			"[postgresGameRepository_DeleteGameByID.Delete] error in the deleting game with id %s into the database.",
			uuid,
		)))
	}

	p.log.Infow(
		fmt.Sprintf(
			"[postgresGameRepository.DeleteGameByID] game with id %s deleted",
			uuid,
		),
		logger.Fields{"Game": uuid},
	)

	return nil
}
