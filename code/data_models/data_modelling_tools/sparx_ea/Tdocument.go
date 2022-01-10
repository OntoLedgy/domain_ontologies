package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TDocument struct {
	DocID       string         `db:"docid"`
	Docname     sql.NullString `db:"docname"`
	Notes       sql.NullString `db:"notes"`
	Style       sql.NullString `db:"style"`
	ElementID   string         `db:"elementid"`
	Elementtype string         `db:"elementtype"`
	Strcontent  sql.NullString `db:"strcontent"`
	Bincontent  sql.NullString `db:"bincontent"`
	Doctype     sql.NullString `db:"doctype"`
	Author      sql.NullString `db:"author"`
	Version     sql.NullString `db:"version"`
	Isactive    sql.NullInt64  `db:"isactive"`
	Sequence    sql.NullInt64  `db:"sequence"`
	Docdate     pg.NullTime    `db:"docdate"`
}
