package sparx_ea

import (
	"database/sql"
)

type TEcf struct {
	EcfID       string          `db:"ecfid"`
	Description sql.NullString  `db:"description"`
	Weight      sql.NullFloat64 `db:"weight"`
	Value       sql.NullFloat64 `db:"value"`
	Notes       sql.NullString  `db:"notes"`
}
