package notes

import (
	"unjuk_keterampilan/app/base"

	"github.com/labstack/echo/v4"
)

type NoteHandler struct {
	noteService NoteServiceImpl
}

type NoteHandlerImpl interface {
	GetNotes(c echo.Context) error
}

func NewNoteHandler(nodeService NoteServiceImpl) *NoteHandler {
	return &NoteHandler{nodeService}
}

func (nh *NoteHandler) GetNotes(c echo.Context) error {
	notes, err := nh.noteService.GetNotes()
	if err != nil {
		return err
	}

	return base.SetResponseSuccess(notes).Thrown(c)
}
