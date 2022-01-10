package sparx_ea

import (
	"database/sql"
)

type TMetrictypes struct {
	Metric        string          `db:"metric"`
	Description   sql.NullString  `db:"description"`
	Numericweight sql.NullFloat64 `db:"numericweight"`
	Notes         sql.NullString  `db:"notes"`
}
