package fxapp

import (
	"time"

	"github.com/zizouhuweidi/retro-station/internal/pkg/config"
	logConfig "github.com/zizouhuweidi/retro-station/internal/pkg/logger/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger/external/fxlog"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger/logrous"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger/models"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger/zap"

	"go.uber.org/fx"
)

func CreateFxApp(
	app *application,
) *fx.App {
	var opts []fx.Option

	opts = append(opts, fx.Provide(app.provides...))

	opts = append(opts, fx.Decorate(app.decorates...))

	opts = append(opts, fx.Invoke(app.invokes...))

	app.options = append(app.options, opts...)

	AppModule := fx.Module("fxapp",
		app.options...,
	)

	var logModule fx.Option
	logOption, err := logConfig.ProvideLogConfig(app.environment)
	if err != nil || logOption == nil {
		logModule = zap.ModuleFunc(app.logger)
	} else if logOption.LogType == models.Logrus {
		logModule = logrous.ModuleFunc(app.logger)
	} else {
		logModule = zap.ModuleFunc(app.logger)
	}

	duration := 30 * time.Second

	// build phase of container will do in this stage, containing provides and invokes but app not started yet and will be started in the future with `fxApp.Register`
	fxApp := fx.New(
		fx.StartTimeout(duration),
		config.ModuleFunc(app.environment),
		logModule,
		fxlog.FxLogger,
		fx.ErrorHook(NewFxErrorHandler(app.logger)),
		AppModule,
	)

	return fxApp
}
