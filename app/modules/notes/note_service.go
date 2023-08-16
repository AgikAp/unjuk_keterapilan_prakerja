package notes

import (
	"unjuk_keterampilan/app/modules/notes/dto"
	"unjuk_keterampilan/app/modules/notes/mapper"
	"unjuk_keterampilan/app/modules/notes/repositorys"
)

type NoteService struct {
	noteRepository repositorys.NoteRepositoryImpl
}

type NoteServiceImpl interface {
	GetNotes() (responses []dto.NoteResponse, err error)
}

func NewNoteService(noteRepository repositorys.NoteRepositoryImpl) *NoteService {
	return &NoteService{noteRepository}
}

func (ns *NoteService) GetNotes() (responses []dto.NoteResponse, err error) {
	notes, err := ns.noteRepository.GetNotes()
	if err != nil {
		return
	}
	responses = mapper.ConvertNoteModelToNoteResponses(notes)
	return
}
