package dto

import "database/sql"

type EditorSubscribeLabelDeleted struct {
	Editor    int            `db:"editor"`
	GID       sql.NullString `db:"gid"`
	DeletedBy int            `db:"deleted_by"`
}
