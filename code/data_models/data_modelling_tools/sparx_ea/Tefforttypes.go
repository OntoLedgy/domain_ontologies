package sparx_ea

import (
	"database/sql"
)

type TEfforttypes struct {
	Efforttype    string          `db:"efforttype"`
	Description   sql.NullString  `db:"description"`
	Numericweight sql.NullFloat64 `db:"numericweight"`
	Notes         sql.NullString  `db:"notes"`
}
