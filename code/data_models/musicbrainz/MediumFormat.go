package dto

import (
	"database/sql"
)

type MediumFormat struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Parent      sql.NullInt64  `db:"parent"`
	ChildOrder  int            `db:"child_order"`
	Year        sql.NullInt64  `db:"year"`
	HasDiscIDs  bool           `db:"has_discids"`
	Description sql.NullString `db:"description"`
	GID         sql.NullString `db:"gid"`
}
