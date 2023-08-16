package notes

import (
	"unjuk_keterampilan/app/modules/notes/dto"
	"unjuk_keterampilan/app/modules/notes/mapper"
	"unjuk_keterampilan/app/modules/notes/repositorys"
	userRepo "unjuk_keterampilan/app/modules/users/repositorys"
)

type NoteService struct {
	noteRepository repositorys.NoteRepositoryImpl
	userRepository userRepo.UserRepositoryImpl
}

type NoteServiceImpl interface {
	GetNotes() (responses []dto.NoteResponse, err error)
	CreateNote(data dto.NoteRequest) (response dto.NoteResponse, err error)
}

func NewNoteService(noteRepository repositorys.NoteRepositoryImpl, userRepository userRepo.UserRepositoryImpl) *NoteService {
	return &NoteService{noteRepository, userRepository}
}

func (ns *NoteService) GetNotes() (responses []dto.NoteResponse, err error) {
	notes, err := ns.noteRepository.GetNotes()
	if err != nil {
		return
	}
	responses = mapper.ConvertNoteModelToNoteResponses(notes)
	return
}

func (ns *NoteService) CreateNote(data dto.NoteRequest) (response dto.NoteResponse, err error) {
	note := mapper.ConvertNoteRequestToNoteModel(data)
	note.User, err = ns.userRepository.GetUserByUsername("system")
	if err != nil {
		return
	}

	note, err = ns.noteRepository.CreateNote(note)
	if err != nil {
		return
	}

	response = mapper.ConvertNoteModelToNoteResponse(note)
	return
}
