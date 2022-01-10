package sparx_ea

import (
	"database/sql"
)

type TObjectproperties struct {
	PropertyID int            `db:"propertyid"`
	ObjectID   sql.NullInt64  `db:"object_id"`
	Property   sql.NullString `db:"property"`
	Value      sql.NullString `db:"value"`
	Notes      sql.NullString `db:"notes"`
	EaGuID     sql.NullString `db:"ea_guid"`
}
