package notes

type NoteHandler struct {
	noteService NoteServiceImpl
}

type NoteHandlerImpl interface{}

func NewNoteHandler(nodeService NoteServiceImpl) *NoteHandler {
	return &NoteHandler{nodeService}
}
