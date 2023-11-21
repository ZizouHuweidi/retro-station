package contracts

import (
	"context"
	"testing"

	postgres "github.com/zizouhuweidi/retro-station/internal/pkg/postgres_pgx"
)

type PostgresPgxContainer interface {
	Start(ctx context.Context, t *testing.T, options ...*PostgresContainerOptions) (*postgres.Pgx, error)
	CreatingContainerOptions(
		ctx context.Context,
		t *testing.T,
		options ...*PostgresContainerOptions,
	) (*postgres.PostgresPgxOptions, error)
	Cleanup(ctx context.Context) error
}
