package notes

import (
	"unjuk_keterampilan/app/base"
	"unjuk_keterampilan/app/modules/notes/dto"

	"github.com/labstack/echo/v4"
)

type NoteHandler struct {
	noteService NoteServiceImpl
}

type NoteHandlerImpl interface {
	GetNotes(c echo.Context) error
	CreateNote(c echo.Context) error
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

func (nh *NoteHandler) CreateNote(c echo.Context) error {
	var data dto.NoteRequest
	if err := c.Bind(&data); err != nil {
		return base.SetResponseError(err.Error()).Thrown(c)
	}

	note, err := nh.noteService.CreateNote(data)
	if err != nil {
		return base.SetResponseError(err.Error()).Thrown(c)
	}

	return base.SetResponseSuccess(note).Thrown(c)
}
