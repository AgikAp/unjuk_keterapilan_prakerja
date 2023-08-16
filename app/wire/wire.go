//go:build wireinject

// + build:wireinject

package wire

import (
	"unjuk_keterampilan/app/routes"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

// var setVariabel = wire.NewSet()

func SetupApp() *echo.Echo {
	wire.Build(
		// setVariabel,
		routes.Routes(),
	)
	return nil
}
