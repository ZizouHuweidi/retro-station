package tracing

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"

	"go.opentelemetry.io/otel/trace"
)

var MessagingTracer trace.Tracer

func init() {
	MessagingTracer = tracing.NewAppTracer(
		"github.com/zizouhuweidi/retro-station/internal/pkg/messaging",
	) // instrumentation name
}
