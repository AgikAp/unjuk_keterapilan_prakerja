package repositorys

import "gorm.io/gorm"

type NoteRepository struct {
	db *gorm.DB
}

type NoteRepositoryImpl interface{}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db}
}
