package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type ReleaseRaw struct {
	ID           int            `db:"id"`
	Title        string         `db:"title"`
	Artist       sql.NullString `db:"artist"`
	Added        pg.NullTime    `db:"added"`
	LastModified pg.NullTime    `db:"last_modified"`
	LookupCount  sql.NullInt64  `db:"lookup_count"`
	ModifyCount  sql.NullInt64  `db:"modify_count"`
	Source       sql.NullInt64  `db:"source"`
	Barcode      sql.NullString `db:"barcode"`
	Comment      string         `db:"comment"`
}
