package model

import "unjuk_keterampilan/app/base"

type User struct {
	base.BaseModel
	Username string
	Email    string
}
