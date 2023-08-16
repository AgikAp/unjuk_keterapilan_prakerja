package dto

import "time"

type NoteResponse struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Note     string    `json:"note"`
	Title    string    `json:"title"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
