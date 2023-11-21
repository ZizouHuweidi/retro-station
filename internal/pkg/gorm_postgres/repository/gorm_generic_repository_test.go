package repository

import (
	"context"
	"log"
	"testing"

	_ "github.com/lib/pq" // postgres driver
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/zizouhuweidi/retro-station/internal/pkg/core/data"
	"github.com/zizouhuweidi/retro-station/internal/pkg/core/data/specification"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	defaultLogger "github.com/zizouhuweidi/retro-station/internal/pkg/logger/default_logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"
	gorm2 "github.com/zizouhuweidi/retro-station/internal/pkg/test/containers/testcontainer/gorm"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
)

// Game is a domain_events entity
type Game struct {
	ID          uuid.UUID
	Name        string
	Weight      int
	IsAvailable bool
}

// GameGorm is DTO used to map Game entity to database
type GameGorm struct {
	ID          uuid.UUID `gorm:"primaryKey;column:id"`
	Name        string    `gorm:"column:name"`
	Weight      int       `gorm:"column:weight"`
	IsAvailable bool      `gorm:"column:is_available"`
}

func init() {
	err := mapper.CreateMap[*GameGorm, *Game]()
	if err != nil {
		log.Fatal(err)
	}

	err = mapper.CreateMap[*Game, *GameGorm]()
	if err != nil {
		log.Fatal(err)
	}
}

func Test_Add(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepository(ctx, t)

	game := &GameGorm{
		ID:          uuid.NewV4(),
		Name:        "added_game",
		Weight:      100,
		IsAvailable: true,
	}

	err = repository.Add(ctx, game)
	if err != nil {
		t.Fatal(err)
	}

	p, err := repository.GetById(ctx, game.ID)
	if err != nil {
		return
	}

	assert.NotNil(t, p)
	assert.Equal(t, game.ID, p.ID)
}

func Test_Add_With_Data_Model(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepositoryWithDataModel(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	game := &Game{
		ID:          uuid.NewV4(),
		Name:        "added_game",
		Weight:      100,
		IsAvailable: true,
	}

	err = repository.Add(ctx, game)
	if err != nil {
		t.Fatal(err)
	}

	p, err := repository.GetById(ctx, game.ID)
	if err != nil {
		return
	}

	assert.NotNil(t, p)
	assert.Equal(t, game.ID, p.ID)
}

func Test_Get_By_Id(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepository(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	all, err := repository.GetAll(ctx, utils.NewListQuery(10, 1))
	if err != nil {
		return
	}
	p := all.Items[0]

	testCases := []struct {
		Name         string
		GameId       uuid.UUID
		ExpectResult *GameGorm
	}{
		{
			Name:         "ExistingGame",
			GameId:       p.ID,
			ExpectResult: p,
		},
		{
			Name:         "NonExistingGame",
			GameId:       uuid.NewV4(),
			ExpectResult: nil,
		},
	}

	for _, c := range testCases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			res, err := repository.GetById(ctx, c.GameId)
			if c.ExpectResult == nil {
				assert.Error(t, err)
				assert.True(t, customErrors.IsNotFoundError(err))
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
				assert.Equal(t, p.ID, res.ID)
			}
		})
	}
}

func Test_Get_By_Id_With_Data_Model(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepositoryWithDataModel(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	all, err := repository.GetAll(ctx, utils.NewListQuery(10, 1))
	if err != nil {
		return
	}
	p := all.Items[0]

	testCases := []struct {
		Name         string
		GameId       uuid.UUID
		ExpectResult *Game
	}{
		{
			Name:         "ExistingGame",
			GameId:       p.ID,
			ExpectResult: p,
		},
		{
			Name:         "NonExistingGame",
			GameId:       uuid.NewV4(),
			ExpectResult: nil,
		},
	}

	for _, c := range testCases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			res, err := repository.GetById(ctx, c.GameId)
			if c.ExpectResult == nil {
				assert.Error(t, err)
				assert.True(t, customErrors.IsNotFoundError(err))
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
				assert.Equal(t, p.ID, res.ID)
			}
		})
	}
}

func Test_Get_All(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepository(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	models, err := repository.GetAll(ctx, utils.NewListQuery(10, 1))
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, models.Items)
}

func Test_Get_All_With_Data_Model(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepositoryWithDataModel(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	models, err := repository.GetAll(ctx, utils.NewListQuery(10, 1))
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, models.Items)
}

func Test_Search(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepository(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	models, err := repository.Search(ctx, "seed_game1", utils.NewListQuery(10, 1))
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, models.Items)
	assert.Equal(t, len(models.Items), 1)
}

func Test_Search_With_Data_Model(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepositoryWithDataModel(ctx, t)

	models, err := repository.Search(ctx, "seed_game1", utils.NewListQuery(10, 1))
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, models.Items)
	assert.Equal(t, len(models.Items), 1)
}

func Test_Where(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepository(ctx, t)

	models, err := repository.GetByFilter(ctx, map[string]interface{}{"name": "seed_game1"})
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, models)
	assert.Equal(t, len(models), 1)
}

func Test_Where_With_Data_Model(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepositoryWithDataModel(ctx, t)

	models, err := repository.GetByFilter(ctx, map[string]interface{}{"name": "seed_game1"})
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, models)
	assert.Equal(t, len(models), 1)
}

func Test_Update(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepository(ctx, t)

	games, err := repository.GetAll(ctx, utils.NewListQuery(10, 1))
	if err != nil {
		t.Fatal(err)
	}
	game := games.Items[0]

	game.Name = "game2_updated"
	err = repository.Update(ctx, game)
	if err != nil {
		t.Fatal(err)
	}

	single, err := repository.GetById(ctx, game.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, single)
	assert.Equal(t, "game2_updated", single.Name)
}

func Test_Update_With_Data_Model(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepositoryWithDataModel(ctx, t)

	games, err := repository.GetAll(ctx, utils.NewListQuery(10, 1))
	if err != nil {
		t.Fatal(err)
	}
	game := games.Items[0]

	game.Name = "game2_updated"
	err = repository.Update(ctx, game)
	if err != nil {
		t.Fatal(err)
	}

	single, err := repository.GetById(ctx, game.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, single)
	assert.Equal(t, "game2_updated", single.Name)
}

func Test_Delete(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepository(ctx, t)

	games, err := repository.GetAll(ctx, utils.NewListQuery(10, 1))
	if err != nil {
		t.Fatal(err)
	}
	game := games.Items[0]

	err = repository.Delete(ctx, game.ID)
	if err != nil {
		return
	}

	single, err := repository.GetById(ctx, game.ID)
	assert.Nil(t, single)
}

func Test_Delete_With_Data_Model(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepositoryWithDataModel(ctx, t)

	games, err := repository.GetAll(ctx, utils.NewListQuery(10, 1))
	if err != nil {
		t.Fatal(err)
	}
	game := games.Items[0]

	err = repository.Delete(ctx, game.ID)
	if err != nil {
		return
	}

	single, err := repository.GetById(ctx, game.ID)
	assert.Nil(t, single)
}

func Test_Count(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepository(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	count := repository.Count(ctx)

	assert.Equal(t, count, int64(2))
}

func Test_Count_With_Data_Model(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepositoryWithDataModel(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	count := repository.Count(ctx)

	assert.Equal(t, count, int64(2))
}

func Test_Find(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepository(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	entities, err := repository.Find(
		ctx,
		specification.And(specification.Equal("is_available", true), specification.Equal("name", "seed_game1")),
	)
	if err != nil {
		return
	}
	assert.Equal(t, len(entities), 1)
}

func Test_Find_With_Data_Model(t *testing.T) {
	ctx := context.Background()
	repository, err := setupGenericGormRepositoryWithDataModel(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	entities, err := repository.Find(
		ctx,
		specification.And(specification.Equal("is_available", true), specification.Equal("name", "seed_game1")),
	)
	if err != nil {
		return
	}
	assert.Equal(t, len(entities), 1)
}

func setupGenericGormRepositoryWithDataModel(
	ctx context.Context,
	t *testing.T,
) (data.GenericRepositoryWithDataModel[*GameGorm, *Game], error) {
	defaultLogger.SetupDefaultLogger()

	db, err := gorm2.NewGormTestContainers(defaultLogger.Logger).Start(ctx, t)
	if err != nil {
		return nil, err
	}

	err = seedAndMigration(ctx, db)
	if err != nil {
		return nil, err
	}

	return NewGenericGormRepositoryWithDataModel[*GameGorm, *Game](db), nil
}

func setupGenericGormRepository(ctx context.Context, t *testing.T) (data.GenericRepository[*GameGorm], error) {
	defaultLogger.SetupDefaultLogger()

	db, err := gorm2.NewGormTestContainers(defaultLogger.Logger).Start(ctx, t)

	err = seedAndMigration(ctx, db)
	if err != nil {
		return nil, err
	}

	return NewGenericGormRepository[*GameGorm](db), nil
}

func seedAndMigration(ctx context.Context, db *gorm.DB) error {
	err := db.AutoMigrate(GameGorm{})
	if err != nil {
		return err
	}

	seedGames := []*GameGorm{
		{
			ID:          uuid.NewV4(),
			Name:        "seed_game1",
			Weight:      100,
			IsAvailable: true,
		},
		{
			ID:          uuid.NewV4(),
			Name:        "seed_game2",
			Weight:      100,
			IsAvailable: true,
		},
	}

	err = db.WithContext(ctx).Create(seedGames).Error
	if err != nil {
		return err
	}
	return nil
}
