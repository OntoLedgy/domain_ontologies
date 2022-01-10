package dto

import (
	"database/sql"
)

type Label struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	ContactInfo sql.NullString `db:"contact_info"`
	Profile     sql.NullString `db:"profile"`
	ParentID    sql.NullInt64  `db:"parent_id"`
	ParentName  sql.NullString `db:"parent_name"`
	DataQuality sql.NullString `db:"data_quality"`
}
