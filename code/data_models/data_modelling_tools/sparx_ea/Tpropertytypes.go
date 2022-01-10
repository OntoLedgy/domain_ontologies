package sparx_ea

import (
	"database/sql"
)

type TPropertytypes struct {
	Property    string         `db:"property"`
	Description sql.NullString `db:"description"`
	Notes       sql.NullString `db:"notes"`
}
