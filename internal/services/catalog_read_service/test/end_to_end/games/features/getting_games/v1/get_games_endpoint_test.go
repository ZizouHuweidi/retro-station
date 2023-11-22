//go:build e2e
// +build e2e

package v1

import (
	"context"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/test_fixture/integration"
)

func TestGetAllGames(t *testing.T) {
	e2eFixture := integration.NewIntegrationTestSharedFixture(t)

	Convey("Get All Games Feature", t, func() {
		e2eFixture.InitializeTest()
		ctx := context.Background()

		Convey("Get all games returns ok status", func() {
			Convey("When a request is made to get all games", func() {
				expect := httpexpect.New(t, e2eFixture.BaseAddress)

				Convey("Then the response status should be OK", func() {
					expect.GET("games").
						WithContext(ctx).
						Expect().
						Status(http.StatusOK)
				})
			})
		})

		e2eFixture.DisposeTest()
	})
}
