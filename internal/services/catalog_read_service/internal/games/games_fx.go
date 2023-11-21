package games

import (
	"github.com/labstack/echo/v4"
	customEcho "github.com/zizouhuweidi/retro-station/internal/pkg/http/custom_echo"
	"github.com/zizouhuweidi/retro-station/internal/pkg/web/route"
	"go.uber.org/fx"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/data/repositories"
	getGameByIdV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/get_game_by_id/v1/endpoints"
	getGamesV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/getting_games/v1/endpoints"
	searchGameV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/searching_games/v1/endpoints"
)

var Module = fx.Module(
	"gamesfx",

	// Other provides
	fx.Provide(repositories.NewRedisGameRepository),
	fx.Provide(repositories.NewMongoGameRepository),

	fx.Provide(fx.Annotate(func(catalogsServer customEcho.EchoHttpServer) *echo.Group {
		var g *echo.Group
		catalogsServer.RouteBuilder().RegisterGroupFunc("/api/v1", func(v1 *echo.Group) {
			group := v1.Group("/games")
			g = group
		})

		return g
	}, fx.ResultTags(`name:"game-echo-group"`))),

	fx.Provide(
		route.AsRoute(getGamesV1.NewGetGamesEndpoint, "game-routes"),
		route.AsRoute(searchGameV1.NewSearchGamesEndpoint, "game-routes"),
		route.AsRoute(getGameByIdV1.NewGetGameByIdEndpoint, "game-routes"),
	),
)
