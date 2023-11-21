package infrastructure

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
)

type InfrastructureConfigurator struct {
	contracts.Application
}

func NewInfrastructureConfigurator(app contracts.Application) *InfrastructureConfigurator {
	return &InfrastructureConfigurator{
		Application: app,
	}
}

func (ic *InfrastructureConfigurator) ConfigInfrastructures() {
	ic.ResolveFunc(func() error {
		return nil
	})
}
