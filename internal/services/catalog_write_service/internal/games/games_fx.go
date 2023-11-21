package games

import (
	"github.com/labstack/echo/v4"
	customEcho "github.com/zizouhuweidi/retro-station/internal/pkg/http/custom_echo"
	"github.com/zizouhuweidi/retro-station/internal/pkg/web/route"
	"go.uber.org/fx"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/data/repositories"
	createGameV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/endpoints"
	deleteGameV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/deleting_game/v1/endpoints"
	getGameByIdV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_game_by_id/v1/endpoints"
	getGamesV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/getting_games/v1/endpoints"
	searchGamesV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/searching_game/v1/endpoints"
	updateGamesV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/updating_game/v1/endpoints"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc"
)

var Module = fx.Module(
	"gamesfx",

	// Other provides
	fx.Provide(repositories.NewPostgresGameRepository),
	fx.Provide(grpc.NewGameGrpcService),

	fx.Provide(fx.Annotate(func(catalogsServer customEcho.EchoHttpServer) *echo.Group {
		var g *echo.Group
		catalogsServer.RouteBuilder().RegisterGroupFunc("/api/v1", func(v1 *echo.Group) {
			group := v1.Group("/games")
			g = group
		})

		return g
	}, fx.ResultTags(`name:"game-echo-group"`))),

	fx.Provide(
		route.AsRoute(createGameV1.NewCreteGameEndpoint, "game-routes"),
		route.AsRoute(updateGamesV1.NewUpdateGameEndpoint, "game-routes"),
		route.AsRoute(getGamesV1.NewGetGamesEndpoint, "game-routes"),
		route.AsRoute(searchGamesV1.NewSearchGamesEndpoint, "game-routes"),
		route.AsRoute(getGameByIdV1.NewGetGameByIdEndpoint, "game-routes"),
		route.AsRoute(deleteGameV1.NewDeleteGameEndpoint, "game-routes"),
	),
)
