package dto

import (
	"database/sql"
)

type WorkAttributeTypeAllowedValue struct {
	ID                int            `db:"id"`
	WorkAttributeType int            `db:"work_attribute_type"`
	Value             sql.NullString `db:"value"`
	Parent            sql.NullInt64  `db:"parent"`
	ChildOrder        int            `db:"child_order"`
	Description       sql.NullString `db:"description"`
	GID               sql.NullString `db:"gid"`
}
