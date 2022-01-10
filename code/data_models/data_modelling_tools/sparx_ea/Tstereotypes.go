package sparx_ea

import (
	"database/sql"
)

type TStereotypes struct {
	Stereotype  string         `db:"stereotype"`
	Appliesto   string         `db:"appliesto"`
	Description sql.NullString `db:"description"`
	Mfenabled   sql.NullInt64  `db:"mfenabled"`
	Mfpath      sql.NullString `db:"mfpath"`
	Metafile    sql.NullString `db:"metafile"`
	Style       sql.NullString `db:"style"`
	EaGuID      sql.NullString `db:"ea_guid"`
	Visualtype  sql.NullString `db:"visualtype"`
}
