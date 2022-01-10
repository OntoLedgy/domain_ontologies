package dto

import (
	"database/sql"
)

type ReleaseGroupSecondaryType struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Parent      sql.NullInt64  `db:"parent"`
	ChildOrder  int            `db:"child_order"`
	Description sql.NullString `db:"description"`
	GID         sql.NullString `db:"gid"`
}
