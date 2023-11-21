package test

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/test"

	"go.uber.org/fx/fxtest"
)

type CatalogsWriteTestApplicationBuilder struct {
	contracts.ApplicationBuilder
	tb fxtest.TB
}

func NewCatalogsWriteTestApplicationBuilder(tb fxtest.TB) *CatalogsWriteTestApplicationBuilder {
	return &CatalogsWriteTestApplicationBuilder{
		ApplicationBuilder: test.NewTestApplicationBuilder(tb),
		tb:                 tb,
	}
}

func (a *CatalogsWriteTestApplicationBuilder) Build() *CatalogsWriteTestApplication {
	return NewCatalogsWriteTestApplication(
		a.tb,
		a.GetProvides(),
		a.GetDecorates(),
		a.Options(),
		a.Logger(),
		a.Environment(),
	)
}
