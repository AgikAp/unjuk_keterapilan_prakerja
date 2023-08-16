package model

import (
	"unjuk_keterampilan/app/base"
	"unjuk_keterampilan/app/modules/users/model"
)

type Note struct {
	base.BaseModel
	UserID string
	Note   string
	Title  string
	User   model.User
}
