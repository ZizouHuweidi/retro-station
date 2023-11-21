package repositories

// https://github.com/Kamva/mgm
// https://github.com/mongodb/mongo-go-driver
// https://blog.logrocket.com/how-to-use-mongodb-with-go/
// https://www.mongodb.com/docs/drivers/go/current/quick-reference/
// https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/
// https://www.mongodb.com/docs

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	uuid2 "github.com/satori/go.uuid"
	"github.com/zizouhuweidi/retro-station/internal/pkg/core/data"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mongodb"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mongodb/repository"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing/attribute"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/models"
)

const (
	gameCollection = "games"
)

type mongoGameRepository struct {
	log                    logger.Logger
	mongoGenericRepository data.GenericRepository[*models.Game]
	tracer                 tracing.AppTracer
}

func NewMongoGameRepository(
	log logger.Logger,
	db *mongo.Client,
	mongoOptions *mongodb.MongoDbOptions,
	tracer tracing.AppTracer,
) data2.gameRepository {
	mongoRepo := repository.NewGenericMongoRepository[*models.Game](
		db,
		mongoOptions.Database,
		gameCollection,
	)
	return &mongoGameRepository{log: log, mongoGenericRepository: mongoRepo, tracer: tracer}
}

func (p *mongoGameRepository) GetAllgames(
	ctx context.Context,
	listQuery *utils.ListQuery,
) (*utils.ListResult[*models.Game], error) {
	ctx, span := p.tracer.Start(ctx, "mongoGameRepository.GetAllgames")
	defer span.End()

	// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/
	result, err := p.mongoGenericRepository.GetAll(ctx, listQuery)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				"[mongoGameRepository_GetAllGames.Paginate] error in the paginate",
			),
		)
	}

	p.log.Infow(
		"[mongoGameRepository.GetAllgames] games loaded",
		logger.Fields{"gamesResult": result},
	)

	span.SetAttributes(attribute.Object("gamesResult", result))

	return result, nil
}

func (p *mongoGameRepository) Searchgames(
	ctx context.Context,
	searchText string,
	listQuery *utils.ListQuery,
) (*utils.ListResult[*models.Game], error) {
	ctx, span := p.tracer.Start(ctx, "mongogameRepository.Searchgames")
	span.SetAttributes(attribute2.String("SearchText", searchText))
	defer span.End()

	result, err := p.mongoGenericRepository.Search(ctx, searchText, listQuery)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				"[mongoGameRepository_Searchgames.Paginate] error in the paginate",
			),
		)
	}

	p.log.Infow(
		fmt.Sprintf(
			"[mongoGameRepository.Searchgames] games loaded for search term '%s'",
			searchText,
		),
		logger.Fields{"gamesResult": result},
	)

	span.SetAttributes(attribute.Object("gamesResult", result))

	return result, nil
}

func (p *mongoGameRepository) GetgameById(
	ctx context.Context,
	uuid string,
) (*models.Game, error) {
	ctx, span := p.tracer.Start(ctx, "mongogameRepository.GetgameById")
	span.SetAttributes(attribute2.String("Id", uuid))
	defer span.End()

	id, err := uuid2.FromString(uuid)
	if err != nil {
		return nil, err
	}

	game, err := p.mongoGenericRepository.GetById(ctx, id)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"[mongogameRepository_GetgameById.FindOne] can't find the game with id %s into the database.",
					uuid,
				),
			),
		)
	}

	span.SetAttributes(attribute.Object("game", game))

	p.log.Infow(
		fmt.Sprintf("[mongogameRepository.GetgameById] game with id %s laoded", uuid),
		logger.Fields{"game": game, "Id": uuid},
	)

	return game, nil
}

func (p *mongoGameRepository) GetgameBygameId(
	ctx context.Context,
	uuid string,
) (*models.Game, error) {
	gameId := uuid
	ctx, span := p.tracer.Start(ctx, "mongogameRepository.GetgameBygameId")
	span.SetAttributes(attribute2.String("gameId", gameId))
	defer span.End()

	game, err := p.mongoGenericRepository.FirstOrDefault(
		ctx,
		map[string]interface{}{"gameId": uuid},
	)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"[mongogameRepository_GetgameById.FindOne] can't find the game with gameId %s into the database.",
					uuid,
				),
			),
		)
	}

	span.SetAttributes(attribute.Object("game", game))

	p.log.Infow(
		fmt.Sprintf(
			"[mongogameRepository.GetgameById] game with gameId %s laoded",
			gameId,
		),
		logger.Fields{"game": game, "gameId": uuid},
	)

	return game, nil
}

func (p *mongoGameRepository) Creategame(
	ctx context.Context,
	game *models.Game,
) (*models.Game, error) {
	ctx, span := p.tracer.Start(ctx, "mongoGameRepository.CreateGame")
	defer span.End()

	err := p.mongoGenericRepository.Add(ctx, game)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				"[mongoGameRepository_CreateGame.InsertOne] error in the inserting game into the database.",
			),
		)
	}

	span.SetAttributes(attribute.Object("game", game))

	p.log.Infow(
		fmt.Sprintf(
			"[mongoGameRepository.CreateGame] game with id '%s' created",
			game.GameId,
		),
		logger.Fields{"game": game, "Id": game.GameId},
	)

	return game, nil
}

func (p *mongoGameRepository) UpdateGame(
	ctx context.Context,
	updateGame *models.Game,
) (*models.Game, error) {
	ctx, span := p.tracer.Start(ctx, "mongoGameRepository.UpdateGame")
	defer span.End()

	err := p.mongoGenericRepository.Update(ctx, updateGame)
	// https://www.mongodb.com/docs/manual/reference/method/db.collection.findOneAndUpdate/
	if err != nil {
		return nil, tracing.TraceErrFromSpan(
			span,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"[mongoGameRepository_UpdateGame.FindOneAndUpdate] error in updating game with id %s into the database.",
					updateGame.GameId,
				),
			),
		)
	}

	span.SetAttributes(attribute.Object("game", updateGame))
	p.log.Infow(
		fmt.Sprintf(
			"[mongoGameRepository.UpdateGame] game with id '%s' updated",
			updateGame.GameId,
		),
		logger.Fields{"game": updateGame, "Id": updateGame.GameId},
	)

	return updateGame, nil
}

func (p *mongoGameRepository) DeleteGameByID(ctx context.Context, uuid string) error {
	ctx, span := p.tracer.Start(ctx, "mongoGameRepository.DeleteGameByID")
	span.SetAttributes(attribute2.String("Id", uuid))
	defer span.End()

	id, err := uuid2.FromString(uuid)
	if err != nil {
		return err
	}

	err = p.mongoGenericRepository.Delete(ctx, id)
	if err != nil {
		return tracing.TraceErrFromSpan(span, errors.WrapIf(err, fmt.Sprintf(
			"[mongoGameRepository_DeleteGameByID.FindOneAndDelete] error in deleting game with id %s from the database.",
			uuid,
		)))
	}

	p.log.Infow(
		fmt.Sprintf("[mongoGameRepository.DeleteGameByID] game with id %s deleted", uuid),
		logger.Fields{"game": uuid},
	)

	return nil
}
