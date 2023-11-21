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

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/params"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/get_game_by_id/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/get_game_by_id/v1/queries"
)

type getGameByIdEndpoint struct {
	params.GameRouteParams
}

func NewGetGameByIdEndpoint(
	params params.GameRouteParams,
) route.Endpoint {
	return &getGameByIdEndpoint{
		GameRouteParams: params,
	}
}

func (ep *getGameByIdEndpoint) MapEndpoint() {
	ep.GamesGroup.GET("/:id", ep.handler())
}

// GetGameByID
// @Tags Games
// @Summary Get game
// @Description Get game by id
// @Accept json
// @Produce json
// @Param id path string true "Game ID"
// @Success 200 {object} dtos.GetGameByIdResponseDto
// @Router /api/v1/games/{id} [get]
func (ep *getGameByIdEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ep.CatalogsMetrics.GetGameByIdHttpRequests.Add(ctx, 1)

		request := &dtos.GetGameByIdRequestDto{}
		if err := c.Bind(request); err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(
				err,
				"[getGameByIdEndpoint_handler.Bind] error in the binding request",
			)
			ep.Logger.Errorf(
				fmt.Sprintf("[getGameByIdEndpoint_handler.Bind] err: %v", badRequestErr),
			)
			return badRequestErr
		}

		query, err := queries.NewGetGameById(request.Id)
		if err != nil {
			validationErr := customErrors.NewValidationErrorWrap(
				err,
				"[getGameByIdEndpoint_handler.StructCtx]  query validation failed",
			)
			ep.Logger.Errorf("[getGameByIdEndpoint_handler.StructCtx] err: {%v}", validationErr)
			return validationErr
		}

		queryResult, err := mediatr.Send[*queries.GetGameById, *dtos.GetGameByIdResponseDto](
			ctx,
			query,
		)
		if err != nil {
			err = errors.WithMessage(
				err,
				"[getGameByIdEndpoint_handler.Send] error in sending GetGameById",
			)
			ep.Logger.Errorw(
				fmt.Sprintf(
					"[getGameByIdEndpoint_handler.Send] id: {%s}, err: {%v}",
					query.Id,
					err,
				),
				logger.Fields{"GameId": query.Id},
			)
			return err
		}

		return c.JSON(http.StatusOK, queryResult)
	}
}
