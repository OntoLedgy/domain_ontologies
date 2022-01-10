package dto

import "database/sql"

type EditorLanguage struct {
	Editor   int            `db:"editor"`
	Language int            `db:"language"`
	Fluency  sql.NullString `db:"fluency"`
}
