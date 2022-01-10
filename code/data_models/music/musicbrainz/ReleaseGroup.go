package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type ReleaseGroup struct {
	ID           int            `db:"id"`
	GID          sql.NullString `db:"gid"`
	Name         string         `db:"name"`
	ArtistCredit int            `db:"artist_credit"`
	Type         sql.NullInt64  `db:"type"`
	Comment      string         `db:"comment"`
	EditsPending int            `db:"edits_pending"`
	LastUpdated  pg.NullTime    `db:"last_updated"`
}
