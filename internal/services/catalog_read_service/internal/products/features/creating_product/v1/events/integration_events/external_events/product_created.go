package externalEvents

import (
	"time"

	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/types"
)

type ProductCreatedV1 struct {
	*types.Message
	ProductId   string    `json:"productId,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}
