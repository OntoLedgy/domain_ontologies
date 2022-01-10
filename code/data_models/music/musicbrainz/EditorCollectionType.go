package dto

import (
	"database/sql"
)

type EditorCollectionType struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	EntityType  string         `db:"entity_type"`
	Parent      sql.NullInt64  `db:"parent"`
	ChildOrder  int            `db:"child_order"`
	Description sql.NullString `db:"description"`
	GID         sql.NullString `db:"gid"`
}
