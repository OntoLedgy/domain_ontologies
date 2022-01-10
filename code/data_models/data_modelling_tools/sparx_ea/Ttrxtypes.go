package sparx_ea

import (
	"database/sql"
)

type TTrxtypes struct {
	Description   sql.NullString  `db:"description"`
	Numericweight sql.NullFloat64 `db:"numericweight"`
	Notes         sql.NullString  `db:"notes"`
	Trx           sql.NullString  `db:"trx"`
	TrxID         int             `db:"trx_id"`
	Style         sql.NullString  `db:"style"`
}
