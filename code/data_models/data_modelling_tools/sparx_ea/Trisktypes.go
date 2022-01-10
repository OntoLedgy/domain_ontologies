package sparx_ea

import (
	"database/sql"
)

type TRisktypes struct {
	Risk          string          `db:"risk"`
	Description   sql.NullString  `db:"description"`
	Numericweight sql.NullFloat64 `db:"numericweight"`
	Notes         sql.NullString  `db:"notes"`
}
