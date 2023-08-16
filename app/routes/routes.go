package routes

import (
	"unjuk_keterampilan/app/modules/notes"

	"github.com/labstack/echo/v4"
)

func Routes(noteHandler notes.NoteHandlerImpl) *echo.Echo {
	e := echo.New()
	api := e.Group("/api")
	note := api.Group("/note")
	note.GET("", noteHandler.GetNotes)
	note.POST("", noteHandler.CreateNote)
	return e
}
