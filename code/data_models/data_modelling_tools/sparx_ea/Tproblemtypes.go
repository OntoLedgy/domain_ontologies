package sparx_ea

import (
	"database/sql"
)

type TProblemtypes struct {
	Problemtype   string          `db:"problemtype"`
	Description   sql.NullString  `db:"description"`
	Numericweight sql.NullFloat64 `db:"numericweight"`
	Notes         sql.NullString  `db:"notes"`
}
