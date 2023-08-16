package model

import "unjuk_keterampilan/app/base"

type Note struct {
	base.BaseModel
	UserID string
	Note   string
	Title  string
}
