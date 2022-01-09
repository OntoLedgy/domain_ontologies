package dto

import (
	"database/sql"
)

type MasterVideo struct {
	ID          int            `db:"id"`
	MasterID    int            `db:"master_id"`
	Duration    sql.NullInt64  `db:"duration"`
	Title       sql.NullString `db:"title"`
	Description sql.NullString `db:"description"`
	Uri         sql.NullString `db:"uri"`
}
