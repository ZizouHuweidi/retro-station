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
	createGameCommand "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/commands"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/dtos"
)

type createGameEndpoint struct {
	params.GameRouteParams
}

func NewCreteGameEndpoint(params params.GameRouteParams) route.Endpoint {
	return &createGameEndpoint{GameRouteParams: params}
}

func (ep *createGameEndpoint) MapEndpoint() {
	ep.GamesGroup.POST("", ep.handler())
}

// CreateGame
// @Tags Games
// @Summary Create game
// @Description Create new game item
// @Accept json
// @Produce json
// @Param CreateGameRequestDto body dtos.CreateGameRequestDto true "Game data"
// @Success 201 {object} dtos.CreateGameResponseDto
// @Router /api/v1/games [post]
func (ep *createGameEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		ep.CatalogsMetrics.CreateGameHttpRequests.Add(ctx, 1)

		request := &dtos.CreateGameRequestDto{}
		if err := c.Bind(request); err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(
				err,
				"[createGameEndpoint_handler.Bind] error in the binding request",
			)
			ep.Logger.Errorf(
				fmt.Sprintf("[createGameEndpoint_handler.Bind] err: %v", badRequestErr),
			)
		}

		command, err := createGameCommand.NewCreateGame(
			request.Name,
			request.Description,
			request.Price,
			request.Genre,
		)
		if err != nil {
			validationErr := customErrors.NewValidationErrorWrap(
				err,
				"[createGameEndpoint_handler.StructCtx] command validation failed",
			)
			ep.Logger.Errorf(
				fmt.Sprintf("[createGameEndpoint_handler.StructCtx] err: {%v}", validationErr),
			)
			return validationErr
		}

		result, err := mediatr.Send[*createGameCommand.CreateGame, *dtos.CreateGameResponseDto](
			ctx,
			command,
		)
		if err != nil {
			err = errors.WithMessage(
				err,
				"[createGameEndpoint_handler.Send] error in sending CreateGame",
			)
			ep.Logger.Errorw(
				fmt.Sprintf(
					"[createGameEndpoint_handler.Send] id: {%s}, err: {%v}",
					command.GameID,
					err,
				),
				logger.Fields{"GameId": command.GameID},
			)
			return err
		}

		return c.JSON(http.StatusCreated, result)
	}
}
