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
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/dtos"
)

type updateGameEndpoint struct {
	params.GameRouteParams
}

func NewUpdateGameEndpoint(
	params params.GameRouteParams,
) route.Endpoint {
	return &updateGameEndpoint{GameRouteParams: params}
}

func (ep *updateGameEndpoint) MapEndpoint() {
	ep.GamesGroup.PUT("/:id", ep.handler())
}

// UpdateGame
// @Tags Games
// @Summary Update game
// @Description Update existing game
// @Accept json
// @Produce json
// @Param UpdateGameRequestDto body dtos.UpdateGameRequestDto true "Game data"
// @Param id path string true "Game ID"
// @Success 204
// @Router /api/v1/games/{id} [put]
func (ep *updateGameEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ep.CatalogsMetrics.UpdateGameHttpRequests.Add(ctx, 1)

		request := &dtos.UpdateGameRequestDto{}
		if err := c.Bind(request); err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(
				err,
				"[updateGameEndpoint_handler.Bind] error in the binding request",
			)
			ep.Logger.Errorf(
				fmt.Sprintf("[updateGameEndpoint_handler.Bind] err: %v", badRequestErr),
			)
			return badRequestErr
		}

		command, err := commands.NewUpdateGame(
			request.GameID,
			request.Name,
			request.Description,
			request.Price,
			request.Genre,
		)
		if err != nil {
			validationErr := customErrors.NewValidationErrorWrap(
				err,
				"[updateGameEndpoint_handler.StructCtx] command validation failed",
			)
			ep.Logger.Errorf(
				fmt.Sprintf("[updateGameEndpoint_handler.StructCtx] err: {%v}", validationErr),
			)
			return validationErr
		}

		_, err = mediatr.Send[*commands.UpdateGame, *mediatr.Unit](ctx, command)
		if err != nil {
			err = errors.WithMessage(
				err,
				"[updateGameEndpoint_handler.Send] error in sending UpdateGame",
			)
			ep.Logger.Errorw(
				fmt.Sprintf(
					"[updateGameEndpoint_handler.Send] id: {%s}, err: {%v}",
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
