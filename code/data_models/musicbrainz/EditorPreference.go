package dto

type EditorPreference struct {
	ID     int    `db:"id"`
	Editor int    `db:"editor"`
	Name   string `db:"name"`
	Value  string `db:"value"`
}
