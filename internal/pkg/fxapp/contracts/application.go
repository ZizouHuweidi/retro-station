package contracts

import (
	"context"

	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"

	"go.uber.org/fx"
)

type Application interface {
	Container
	RegisterHook(function interface{})
	Run()
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Wait() <-chan fx.ShutdownSignal
	Logger() logger.Logger
	Environment() environemnt.Environment
}
