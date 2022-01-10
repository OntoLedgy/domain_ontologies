package sparx_ea

import (
	"database/sql"
)

type TTcf struct {
	TcfID       string          `db:"tcfid"`
	Description sql.NullString  `db:"description"`
	Weight      sql.NullFloat64 `db:"weight"`
	Value       sql.NullFloat64 `db:"value"`
	Notes       sql.NullString  `db:"notes"`
}
