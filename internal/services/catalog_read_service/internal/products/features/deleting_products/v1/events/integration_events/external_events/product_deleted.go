package externalEvents

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
)

type ProductDeletedV1 struct {
	*types.Message
	ProductId string `json:"productId,omitempty"`
}
