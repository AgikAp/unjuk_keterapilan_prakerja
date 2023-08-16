package mapper

import (
	"unjuk_keterampilan/app/modules/notes/dto"
	"unjuk_keterampilan/app/modules/notes/model"
)

func ConvertNoteModelToNoteResponse(noteModel model.Note) (response dto.NoteResponse) {
	response.ID = noteModel.ID
	response.Note = noteModel.Note
	response.Title = noteModel.Title
	response.Username = noteModel.User.Username
	response.Created = noteModel.CreatedAt
	response.Updated = noteModel.UpdatedAt
	return
}

func ConvertNoteModelToNoteResponses(noteModels []model.Note) (responses []dto.NoteResponse) {
	for _, note := range noteModels {
		responses = append(responses, ConvertNoteModelToNoteResponse(note))
	}
	return
}

func ConvertNoteRequestToNoteModel(noteRequest dto.NoteRequest) (response model.Note) {
	response.Title = noteRequest.Title
	response.Note = noteRequest.Notes
	return
}
