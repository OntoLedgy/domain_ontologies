package dto

type EditorSubscribeArtist struct {
	ID           int `db:"id"`
	Editor       int `db:"editor"`
	Artist       int `db:"artist"`
	LastEditSent int `db:"last_edit_sent"`
}
