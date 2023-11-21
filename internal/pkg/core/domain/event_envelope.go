package domain

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/core/metadata"
)

type EventEnvelope struct {
	EventData interface{}
	Metadata  metadata.Metadata
}
