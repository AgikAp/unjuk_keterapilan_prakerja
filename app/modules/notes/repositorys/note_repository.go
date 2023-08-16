package repositorys

import (
	"unjuk_keterampilan/app/modules/notes/model"

	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
}

type NoteRepositoryImpl interface {
	GetNotes() (response []model.Note, err error)
	CreateNote(data model.Note) (response model.Note, err error)
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db}
}

func (nr *NoteRepository) GetNotes() (response []model.Note, err error) {
	err = nr.db.Model(model.Note{}).Preload("User").Find(&response).Error
	return
}

func (nr *NoteRepository) CreateNote(data model.Note) (response model.Note, err error) {
	err = nr.db.Model(model.Note{}).Create(&data).Where("id = ?", data.ID).First(&response).Error
	return
}
