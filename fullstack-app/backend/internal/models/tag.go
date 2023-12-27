package models

type Tag struct {
	Id     int    `json:"id"`
	NoteID int    `json:"note_id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}
