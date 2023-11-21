package defaultLogger

import (
	"os"

	"github.com/zizouhuweidi/retro-station/internal/pkg/constants"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger/logrous"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger/models"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger/zap"
)

var Logger logger.Logger

func SetupDefaultLogger() {
	logType := os.Getenv("LogConfig_LogType")

	switch logType {
	case "Zap", "":
		Logger = zap.NewZapLogger(
			&config.LogOptions{LogType: models.Zap, CallerEnabled: false},
			constants.Dev,
		)
		break
	case "Logrus":
		Logger = logrous.NewLogrusLogger(
			&config.LogOptions{LogType: models.Logrus, CallerEnabled: false},
			constants.Dev,
		)
		break
	default:
	}
}
