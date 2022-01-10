package sparx_ea

import (
	"database/sql"
)

type TMainttypes struct {
	Mainttype     string          `db:"mainttype"`
	Description   sql.NullString  `db:"description"`
	Numericweight sql.NullFloat64 `db:"numericweight"`
	Notes         sql.NullString  `db:"notes"`
}
