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

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/params"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/searching_games/v1/dtos"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/searching_games/v1/queries"
)

type searchGamesEndpoint struct {
	params.GameRouteParams
}

func NewSearchGamesEndpoint(
	params params.GameRouteParams,
) route.Endpoint {
	return &searchGamesEndpoint{
		GameRouteParams: params,
	}
}

func (ep *searchGamesEndpoint) MapEndpoint() {
	ep.GamesGroup.GET("/search", ep.handler())
}

// SearchGames
// @Tags Games
// @Summary Search games
// @Description Search games
// @Accept json
// @Produce json
// @Param searchGamesRequestDto query dtos.SearchGamesRequestDto false "SearchGamesRequestDto"
// @Success 200 {object} dtos.SearchGamesResponseDto
// @Router /api/v1/games/search [get]
func (ep *searchGamesEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ep.CatalogsMetrics.SearchGameHttpRequests.Add(ctx, 1)

		listQuery, err := utils.GetListQueryFromCtx(c)
		if err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(
				err,
				"[searchGamesEndpoint_handler.GetListQueryFromCtx] error in getting data from query string",
			)
			ep.Logger.Errorf(
				fmt.Sprintf(
					"[searchGamesEndpoint_handler.GetListQueryFromCtx] err: %v",
					badRequestErr,
				),
			)
			return err
		}

		request := &dtos.SearchGamesRequestDto{ListQuery: listQuery}

		if err := c.Bind(request); err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(
				err,
				"[searchGamesEndpoint_handler.Bind] error in the binding request",
			)
			ep.Logger.Errorf(
				fmt.Sprintf("[searchGamesEndpoint_handler.Bind] err: %v", badRequestErr),
			)
			return badRequestErr
		}

		query := &queries.SearchGames{
			SearchText: request.SearchText,
			ListQuery:  request.ListQuery,
		}

		if err := query.Validate(); err != nil {
			validationErr := customErrors.NewValidationErrorWrap(
				err,
				"[searchGamesEndpoint_handler.StructCtx]  query validation failed",
			)
			ep.Logger.Errorf("[searchGamesEndpoint_handler.StructCtx] err: {%v}", validationErr)
			return validationErr
		}

		queryResult, err := mediatr.Send[*queries.SearchGames, *dtos.SearchGamesResponseDto](
			ctx,
			query,
		)
		if err != nil {
			err = errors.WithMessage(
				err,
				"[searchGamesEndpoint_handler.Send] error in sending SearchGames",
			)
			ep.Logger.Error(fmt.Sprintf("[searchGamesEndpoint_handler.Send] err: {%v}", err))
			return err
		}

		return c.JSON(http.StatusOK, queryResult)
	}
}
