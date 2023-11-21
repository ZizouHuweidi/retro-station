package core

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/core/serializer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/core/serializer/json"
	defaultLogger "github.com/zizouhuweidi/retro-station/internal/pkg/logger/default_logger"

	"go.uber.org/fx"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module(
	"corefx",
	fx.Provide(
		json.NewDefaultSerializer,
		serializer.NewDefaultEventSerializer,
		serializer.NewDefaultMetadataSerializer,
	),
	fx.Invoke(defaultLogger.SetupDefaultLogger),
)
