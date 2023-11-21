package app

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
)

type CatalogsWriteApplicationBuilder struct {
	contracts.ApplicationBuilder
}

func NewCatalogsWriteApplicationBuilder() *CatalogsWriteApplicationBuilder {
	builder := &CatalogsWriteApplicationBuilder{fxapp.NewApplicationBuilder()}

	return builder
}

func (a *CatalogsWriteApplicationBuilder) Build() *CatalogsWriteApplication {
	return NewCatalogsWriteApplication(
		a.GetProvides(),
		a.GetDecorates(),
		a.Options(),
		a.Logger(),
		a.Environment(),
	)
}
