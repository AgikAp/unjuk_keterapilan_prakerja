package repositorys

import (
	"unjuk_keterampilan/app/modules/users/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryImpl interface {
	GetUserByUsername(id string) (response model.User, err error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserByUsername(id string) (response model.User, err error) {
	err = ur.db.Model(model.User{}).Where("username = ?", id).First(&response).Error
	return
}
