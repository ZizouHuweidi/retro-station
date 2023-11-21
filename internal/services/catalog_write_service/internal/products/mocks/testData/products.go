package testData

import (
	"time"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/models"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
)

var Products = []*models.Product{
	{
		ProductId:   uuid.NewV4(),
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       gofakeit.Price(100, 1000),
	},
	{
		ProductId:   uuid.NewV4(),
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       gofakeit.Price(100, 1000),
	},
}
