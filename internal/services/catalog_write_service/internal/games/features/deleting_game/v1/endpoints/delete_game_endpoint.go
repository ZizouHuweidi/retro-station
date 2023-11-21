package endpoints

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/web/route"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/params"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/deleting_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/deleting_game/v1/dtos"
)

type deleteGameEndpoint struct {
	params.GameRouteParams
}

func NewDeleteGameEndpoint(
	params params.GameRouteParams,
) route.Endpoint {
	return &deleteGameEndpoint{GameRouteParams: params}
}

func (ep *deleteGameEndpoint) MapEndpoint() {
	ep.GamesGroup.DELETE("/:id", ep.handler())
}

// DeleteGame
// @Tags Games
// @Summary Delete game
// @Description Delete existing game
// @Accept json
// @Produce json
// @Success 204
// @Param id path string true "Game ID"
// @Router /api/v1/games/{id} [delete]
func (ep *deleteGameEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ep.CatalogsMetrics.DeleteGameHttpRequests.Add(ctx, 1)

		request := &dtos.DeleteGameRequestDto{}
		if err := c.Bind(request); err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(
				err,
				"[deleteGameEndpoint_handler.Bind] error in the binding request",
			)
			ep.Logger.Errorf(
				fmt.Sprintf("[deleteGameEndpoint_handler.Bind] err: %v", badRequestErr),
			)
			return badRequestErr
		}

		command, err := commands.NewDeleteGame(request.GameID)
		if err != nil {
			validationErr := customErrors.NewValidationErrorWrap(
				err,
				"[deleteGameEndpoint_handler.StructCtx] command validation failed",
			)
			ep.Logger.Errorf(
				fmt.Sprintf("[deleteGameEndpoint_handler.StructCtx] err: {%v}", validationErr),
			)
			return validationErr
		}

		_, err = mediatr.Send[*commands.DeleteGame, *mediatr.Unit](ctx, command)

		if err != nil {
			err = errors.WithMessage(
				err,
				"[deleteGameEndpoint_handler.Send] error in sending DeleteGame",
			)
			ep.Logger.Errorw(
				fmt.Sprintf(
					"[deleteGameEndpoint_handler.Send] id: {%s}, err: {%v}",
					command.GameID,
					err,
				),
				logger.Fields{"GameId": command.GameID},
			)
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}
}
