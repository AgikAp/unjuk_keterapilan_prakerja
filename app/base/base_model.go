package base

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"<-create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (bm *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if bm.ID == "" {
		bm.ID = uuid.New().String()
	}

	return nil
}
