package sparx_ea

import (
	"database/sql"
)

type TScenariotypes struct {
	Scenariotype  string          `db:"scenariotype"`
	Description   sql.NullString  `db:"description"`
	Numericweight sql.NullFloat64 `db:"numericweight"`
	Notes         sql.NullString  `db:"notes"`
}
