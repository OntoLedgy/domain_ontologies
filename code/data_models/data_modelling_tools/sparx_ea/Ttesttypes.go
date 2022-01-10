package sparx_ea

import (
	"database/sql"
)

type TTesttypes struct {
	Testtype      string          `db:"testtype"`
	Description   sql.NullString  `db:"description"`
	Numericweight sql.NullFloat64 `db:"numericweight"`
	Notes         sql.NullString  `db:"notes"`
}
