package sparx_ea

import (
	"database/sql"
)

type TOperationtag struct {
	PropertyID int            `db:"propertyid"`
	ElementID  sql.NullInt64  `db:"elementid"`
	Property   sql.NullString `db:"property"`
	Value      sql.NullString `db:"value"`
	Notes      sql.NullString `db:"notes"`
	EaGuID     sql.NullString `db:"ea_guid"`
}
