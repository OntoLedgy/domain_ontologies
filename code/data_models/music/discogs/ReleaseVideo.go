package dto

import (
	"database/sql"
)

type ReleaseVideo struct {
	ID          int            `db:"id"`
	ReleaseID   int            `db:"release_id"`
	Duration    sql.NullInt64  `db:"duration"`
	Title       sql.NullString `db:"title"`
	Description sql.NullString `db:"description"`
	Uri         sql.NullString `db:"uri"`
}
