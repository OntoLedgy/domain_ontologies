package dto

import (
	"database/sql"
)

type ReleaseLabel struct {
	ID        int            `db:"id"`
	ReleaseID int            `db:"release_id"`
	LabelID   sql.NullInt64  `db:"label_id"`
	LabelName string         `db:"label_name"`
	Catno     sql.NullString `db:"catno"`
}
