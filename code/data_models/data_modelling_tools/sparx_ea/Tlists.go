package sparx_ea

import (
	"database/sql"
)

type TLists struct {
	ListID   string         `db:"listid"`
	Category string         `db:"category"`
	Name     string         `db:"name"`
	Nval     sql.NullInt64  `db:"nval"`
	Notes    sql.NullString `db:"notes"`
}
