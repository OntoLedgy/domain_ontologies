package sparx_ea

import (
	"database/sql"
)

type TAuthors struct {
	Authorname string         `db:"authorname"`
	Roles      sql.NullString `db:"roles"`
	Notes      sql.NullString `db:"notes"`
}
