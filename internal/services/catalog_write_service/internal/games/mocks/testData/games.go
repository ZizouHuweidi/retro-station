package testData

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/models"
)

var Games = []*models.Game{
	{
		GameId:      uuid.NewV4(),
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       gofakeit.Price(100, 1000),
		Genre:       gofakeit.Adjective(),
	},
	{
		GameId:      uuid.NewV4(),
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       gofakeit.Price(100, 1000),
		Genre:       gofakeit.Adjective(),
	},
}
