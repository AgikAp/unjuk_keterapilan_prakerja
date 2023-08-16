//go:build wireinject

// + build:wireinject

package wire

import (
	"unjuk_keterampilan/app/databases"
	"unjuk_keterampilan/app/modules/notes"
	"unjuk_keterampilan/app/modules/notes/repositorys"
	"unjuk_keterampilan/app/routes"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var setVariabel = wire.NewSet(
	repositorys.NewNoteRepository,
	wire.Bind(new(repositorys.NoteRepositoryImpl), new(*repositorys.NoteRepository)),

	notes.NewNoteService,
	wire.Bind(new(notes.NoteServiceImpl), new(*notes.NoteService)),

	notes.NewNoteHandler,
	wire.Bind(new(notes.NoteHandlerImpl), new(*notes.NoteHandler)),
)

func SetupApp() *echo.Echo {
	wire.Build(
		databases.NewConnection,
		setVariabel,
		routes.Routes,
	)
	return nil
}
