package sparx_ea

import (
	"database/sql"
)

type TSnapshot struct {
	SnapshotID   string         `db:"snapshotid"`
	SeriesID     string         `db:"seriesid"`
	Position     sql.NullInt64  `db:"position"`
	Snapshotname sql.NullString `db:"snapshotname"`
	Notes        sql.NullString `db:"notes"`
	Style        sql.NullString `db:"style"`
	ElementID    string         `db:"elementid"`
	Elementtype  string         `db:"elementtype"`
	Strcontent   sql.NullString `db:"strcontent"`
	Bincontent1  sql.NullString `db:"bincontent1"`
	Bincontent2  sql.NullString `db:"bincontent2"`
}
