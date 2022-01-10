package sparx_ea

import (
	"database/sql"
)

type TRequiretypes struct {
	Requirement   string          `db:"requirement"`
	Description   sql.NullString  `db:"description"`
	Numericweight sql.NullFloat64 `db:"numericweight"`
	Notes         sql.NullString  `db:"notes"`
}
