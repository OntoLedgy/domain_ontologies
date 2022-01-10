package sparx_ea

import (
	"database/sql"
)

type TPaletteitem struct {
	PaletteID sql.NullInt64 `db:"paletteid"`
	ItemID    sql.NullInt64 `db:"itemid"`
}
