package sparx_ea

import (
	"database/sql"
)

type TAttribute struct {
	ObjectID        sql.NullInt64  `db:"object_id"`
	Name            sql.NullString `db:"name"`
	Scope           sql.NullString `db:"scope"`
	Stereotype      sql.NullString `db:"stereotype"`
	Containment     sql.NullString `db:"containment"`
	Isstatic        sql.NullInt64  `db:"isstatic"`
	Iscollection    sql.NullInt64  `db:"iscollection"`
	Isordered       sql.NullInt64  `db:"isordered"`
	Allowduplicates sql.NullInt64  `db:"allowduplicates"`
	Lowerbound      sql.NullString `db:"lowerbound"`
	Upperbound      sql.NullString `db:"upperbound"`
	Container       sql.NullString `db:"container"`
	Notes           sql.NullString `db:"notes"`
	Derived         sql.NullString `db:"derived"`
	ID              int            `db:"id"`
	Pos             sql.NullInt64  `db:"pos"`
	Genoption       sql.NullString `db:"genoption"`
	Length          sql.NullInt64  `db:"length"`
	Precision       sql.NullInt64  `db:"precision"`
	Scale           sql.NullInt64  `db:"scale"`
	Const           sql.NullInt64  `db:"const"`
	Style           sql.NullString `db:"style"`
	Classifier      sql.NullString `db:"classifier"`
	Default         sql.NullString `db:"Default"`
	Type            sql.NullString `db:"type"`
	EaGuID          sql.NullString `db:"ea_guid"`
	Styleex         sql.NullString `db:"styleex"`
}
