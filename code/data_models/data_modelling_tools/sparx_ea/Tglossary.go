package sparx_ea

import (
	"database/sql"
)

type TGlossary struct {
	Term       sql.NullString `db:"term"`
	Type       sql.NullString `db:"type"`
	Meaning    sql.NullString `db:"meaning"`
	GlossaryID int            `db:"glossaryid"`
}
