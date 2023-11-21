package postgrespxg

import (
	"context"
	"testing"

	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	postgres "github.com/zizouhuweidi/retro-station/internal/pkg/postgres_pgx"
)

var PostgresPgxContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *postgres.PostgresPgxOptions, logger logger.Logger) (*postgres.PostgresPgxOptions, error) {
		return NewPostgresPgxContainers(logger).CreatingContainerOptions(ctx, t)
	}
}
