package dto

import (
	"database/sql"
)

type PlaceAttributeTypeAllowedValue struct {
	ID                 int            `db:"id"`
	PlaceAttributeType int            `db:"place_attribute_type"`
	Value              sql.NullString `db:"value"`
	Parent             sql.NullInt64  `db:"parent"`
	ChildOrder         int            `db:"child_order"`
	Description        sql.NullString `db:"description"`
	GID                sql.NullString `db:"gid"`
}
