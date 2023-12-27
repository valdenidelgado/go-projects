package models

type Link struct {
	Id        int    `json:"id"`
	NoteID    int    `json:"note_id"`
	Url       string `json:"url"`
	CreatedAt string `json:"created_at"`
}
