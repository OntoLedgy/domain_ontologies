package sparx_ea

import (
	"database/sql"
)

type TComplexitytypes struct {
	Complexity    string        `db:"complexity"`
	Numericweight sql.NullInt64 `db:"numericweight"`
}
