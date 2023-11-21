package gorm

import (
	"context"
	"testing"

	gormPostgres "github.com/zizouhuweidi/retro-station/internal/pkg/gorm_postgres"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
)

var GormContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *gormPostgres.GormOptions, logger logger.Logger) (*gormPostgres.GormOptions, error) {
		return NewGormTestContainers(logger).CreatingContainerOptions(ctx, t)
	}
}
