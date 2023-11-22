//go:build unit
// +build unit

package games

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mapper"

	dtoV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/dto/v1"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/test_fixtures/unit_test"
)

type mappingProfileUnitTests struct {
	*unit_test.UnitTestSharedFixture
}

func TestMappingProfileUnit(t *testing.T) {
	suite.Run(
		t,
		&mappingProfileUnitTests{UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t)},
	)
}

func (m *mappingProfileUnitTests) Test_Mappings() {
	gameModel := &models.Game{
		GameId:      uuid.NewV4(),
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.EmojiDescription(),
		Price:       gofakeit.Price(100, 1000),
	}

	gameDto := &dtoV1.GameDto{
		GameId:      uuid.NewV4(),
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.EmojiDescription(),
		Price:       gofakeit.Price(100, 1000),
	}

	m.Run("Should_Map_Game_To_GameDto", func() {
		d, err := mapper.Map[*dtoV1.GameDto](gameModel)
		m.Require().NoError(err)
		m.Equal(gameModel.GameId, d.GameId)
		m.Equal(gameModel.Name, d.Name)
	})

	m.Run("Should_Map_Nil_Game_To_GameDto", func() {
		d, err := mapper.Map[*dtoV1.GameDto](*new(models.Game))
		m.Require().NoError(err)
		m.Nil(d)
	})

	m.Run("Should_Map_GameDto_To_Game", func() {
		d, err := mapper.Map[*models.Game](gameDto)
		m.Require().NoError(err)
		m.Equal(gameDto.GameId, d.GameId)
		m.Equal(gameDto.Name, d.Name)
	})

	m.Run("Should_Map_Nil_GameDto_To_Game", func() {
		d, err := mapper.Map[*models.Game](*new(dtoV1.GameDto))
		m.Require().NoError(err)
		m.Nil(d)
	})
}
