package dto

type EditorSubscribeLabel struct {
	ID           int `db:"id"`
	Editor       int `db:"editor"`
	Label        int `db:"label"`
	LastEditSent int `db:"last_edit_sent"`
}
