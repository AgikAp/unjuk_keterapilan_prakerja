package notes

import "unjuk_keterampilan/app/modules/notes/repositorys"

type NoteService struct {
	noteRepository repositorys.NoteRepositoryImpl
}

type NoteServiceImpl interface{}

func NewNoteRepository(noteRepository repositorys.NoteRepositoryImpl) *NoteService {
	return &NoteService{noteRepository}
}
