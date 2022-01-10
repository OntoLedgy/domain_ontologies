package sparx_ea

import (
	"database/sql"
)

type TCategory struct {
	CategoryID int            `db:"categoryid"`
	Name       sql.NullString `db:"name"`
	Type       sql.NullString `db:"type"`
	Notes      sql.NullString `db:"notes"`
}
