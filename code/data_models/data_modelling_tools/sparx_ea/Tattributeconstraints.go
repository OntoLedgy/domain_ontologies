package sparx_ea

import (
	"database/sql"
)

type TAttributeconstraints struct {
	ObjectID   sql.NullInt64  `db:"object_id"`
	Constraint string         `db:"Constraint"`
	Attname    sql.NullString `db:"attname"`
	Type       sql.NullString `db:"type"`
	Notes      sql.NullString `db:"notes"`
	ID         int            `db:"id"`
}
