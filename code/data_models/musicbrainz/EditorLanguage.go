package dto

type EditorLanguage struct {
	Editor   int            `db:"editor"`
	Language int            `db:"language"`
	Fluency  sql.NullString `db:"fluency"`
}
