package dto

import (
	"database/sql"
)

type LabelImage struct {
	LabelID int            `db:"label_id"`
	Type    sql.NullString `db:"type"`
	WIDth   sql.NullInt64  `db:"width"`
	Height  sql.NullInt64  `db:"height"`
}
