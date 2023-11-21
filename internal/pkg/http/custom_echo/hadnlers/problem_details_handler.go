package customHadnlers

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/problemDetails"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	errorUtils "github.com/zizouhuweidi/retro-station/internal/pkg/utils/error_utils"

	"github.com/labstack/echo/v4"
)

func ProblemHandlerFunc(err error, c echo.Context, logger logger.Logger) {
	prb := problemDetails.ParseError(err)

	if prb != nil {
		if !c.Response().Committed {
			if _, err := problemDetails.WriteTo(prb, c.Response()); err != nil {
				logger.Error(err)
			}
		}
	} else {
		if !c.Response().Committed {
			prb := problemDetails.NewInternalServerProblemDetail(err.Error(), errorUtils.ErrorsWithStack(err))
			if _, err := problemDetails.WriteTo(prb, c.Response()); err != nil {
				logger.Error(err)
			}
		}
	}
}
