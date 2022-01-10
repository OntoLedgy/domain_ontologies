package sparx_ea

import (
	"database/sql"
)

type TPalette struct {
	PaletteID int            `db:"paletteid"`
	Name      sql.NullString `db:"name"`
	Type      sql.NullString `db:"type"`
}
