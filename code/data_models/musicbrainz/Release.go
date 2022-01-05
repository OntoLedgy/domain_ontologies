package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Release struct {
	ID           int            `db:"id"`
	GID          sql.NullString `db:"gid"`
	Name         string         `db:"name"`
	ArtistCredit int            `db:"artist_credit"`
	ReleaseGroup int            `db:"release_group"`
	Status       sql.NullInt64  `db:"status"`
	Packaging    sql.NullInt64  `db:"packaging"`
	Language     sql.NullInt64  `db:"language"`
	Script       sql.NullInt64  `db:"script"`
	Barcode      sql.NullString `db:"barcode"`
	Comment      string         `db:"comment"`
	EditsPending int            `db:"edits_pending"`
	Quality      int            `db:"quality"`
	LastUpdated  pg.NullTime    `db:"last_updated"`
}
