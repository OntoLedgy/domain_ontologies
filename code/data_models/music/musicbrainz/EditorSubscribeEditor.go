package dto

type EditorSubscribeEditor struct {
	ID               int `db:"id"`
	Editor           int `db:"editor"`
	SubscribedEditor int `db:"subscribed_editor"`
	LastEditSent     int `db:"last_edit_sent"`
}
