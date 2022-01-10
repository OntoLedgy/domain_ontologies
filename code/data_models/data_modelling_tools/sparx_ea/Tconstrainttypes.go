package sparx_ea

import (
	"database/sql"
)

type TConstrainttypes struct {
	Constraint  string         `db:"Constraint"`
	Description sql.NullString `db:"description"`
	Notes       sql.NullString `db:"notes"`
}
