package sparx_ea

import (
	"database/sql"
)

type TTaggedvalue struct {
	PropertyID string         `db:"propertyid"`
	ElementID  string         `db:"elementid"`
	Baseclass  string         `db:"baseclass"`
	Tagvalue   sql.NullString `db:"tagvalue"`
	Notes      sql.NullString `db:"notes"`
}
