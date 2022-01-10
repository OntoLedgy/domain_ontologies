package dto

import (
	"database/sql"
)

type ReleaseAttributeTypeAllowedValue struct {
	ID                   int            `db:"id"`
	ReleaseAttributeType int            `db:"release_attribute_type"`
	Value                sql.NullString `db:"value"`
	Parent               sql.NullInt64  `db:"parent"`
	ChildOrder           int            `db:"child_order"`
	Description          sql.NullString `db:"description"`
	GID                  sql.NullString `db:"gid"`
}
