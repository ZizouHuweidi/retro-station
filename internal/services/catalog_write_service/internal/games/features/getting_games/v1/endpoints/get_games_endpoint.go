package endpoints

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"
	"github.com/zizouhuweidi/retro-station/internal/pkg/utils"
	"github.com/zizouhuweidi/retro-station/internal/pkg/web/route"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/params"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_games/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_games/v1/queries"
)

type getGamesEndpoint struct {
	params.GameRouteParams
}

func NewGetGamesEndpoint(
	params params.GameRouteParams,
) route.Endpoint {
	return &getGamesEndpoint{GameRouteParams: params}
}

func (ep *getGamesEndpoint) MapEndpoint() {
	ep.GamesGroup.GET("", ep.handler())
}

// GetAllGames
// @Tags Games
// @Summary Get all game
// @Description Get all games
// @Accept json
// @Produce json
// @Param getGamesRequestDto query dtos.GetGamesRequestDto false "GetGamesRequestDto"
// @Success 200 {object} dtos.GetGamesResponseDto
// @Router /api/v1/games [get]
func (ep *getGamesEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ep.CatalogsMetrics.GetGamesHttpRequests.Add(ctx, 1)

		listQuery, err := utils.GetListQueryFromCtx(c)
		if err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(
				err,
				"[getGamesEndpoint_handler.GetListQueryFromCtx] error in getting data from query string",
			)
			ep.Logger.Errorf(
				fmt.Sprintf(
					"[getGamesEndpoint_handler.GetListQueryFromCtx] err: %v",
					badRequestErr,
				),
			)
			return err
		}

		request := &dtos.GetGamesRequestDto{ListQuery: listQuery}
		if err := c.Bind(request); err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(
				err,
				"[getGamesEndpoint_handler.Bind] error in the binding request",
			)
			ep.Logger.Errorf(
				fmt.Sprintf("[getGamesEndpoint_handler.Bind] err: %v", badRequestErr),
			)
			return badRequestErr
		}

		query, err := queries.NewGetGames(request.ListQuery)
		if err != nil {
			return err
		}

		queryResult, err := mediatr.Send[*queries.GetGames, *dtos.GetGamesResponseDto](
			ctx,
			query,
		)
		if err != nil {
			err = errors.WithMessage(
				err,
				"[getGamesEndpoint_handler.Send] error in sending GetGames",
			)
			ep.Logger.Error(fmt.Sprintf("[getGamesEndpoint_handler.Send] err: {%v}", err))
			return err
		}

		return c.JSON(http.StatusOK, queryResult)
	}
}
