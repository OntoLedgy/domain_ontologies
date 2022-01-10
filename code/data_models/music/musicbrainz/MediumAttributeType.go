package dto

import (
	"database/sql"
)

type MediumAttributeType struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Comment     string         `db:"comment"`
	FreeText    bool           `db:"free_text"`
	Parent      sql.NullInt64  `db:"parent"`
	ChildOrder  int            `db:"child_order"`
	Description sql.NullString `db:"description"`
	GID         sql.NullString `db:"gid"`
}
