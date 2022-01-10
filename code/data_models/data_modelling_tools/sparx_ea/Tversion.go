package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TVersion struct {
	ElementID    string         `db:"elementid"`
	VersionID    string         `db:"versionid"`
	Elementtype  string         `db:"elementtype"`
	Flags        sql.NullString `db:"flags"`
	Externalfile sql.NullString `db:"externalfile"`
	Notes        sql.NullString `db:"notes"`
	Owner        sql.NullString `db:"owner"`
	Versiondate  pg.NullTime    `db:"versiondate"`
	Branch       sql.NullString `db:"branch"`
	ElementXML   sql.NullString `db:"elementxml"`
}
