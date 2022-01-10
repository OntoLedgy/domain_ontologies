package dto

type EditorSubscribeSeries struct {
	ID           int `db:"id"`
	Editor       int `db:"editor"`
	Series       int `db:"series"`
	LastEditSent int `db:"last_edit_sent"`
}
