package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Recording struct {
	ID           int            `db:"id"`
	GID          sql.NullString `db:"gid"`
	Name         string         `db:"name"`
	ArtistCredit int            `db:"artist_credit"`
	Length       sql.NullInt64  `db:"length"`
	Comment      string         `db:"comment"`
	EditsPending int            `db:"edits_pending"`
	LastUpdated  pg.NullTime    `db:"last_updated"`
	VIDeo        bool           `db:"video"`
}
