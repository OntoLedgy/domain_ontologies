package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type ArtistCredit struct {
	ID           int           `db:"id"`
	Name         string        `db:"name"`
	ArtistCount  int           `db:"artist_count"`
	RefCount     sql.NullInt64 `db:"ref_count"`
	Created      pg.NullTime   `db:"created"`
	EditsPending int           `db:"edits_pending"`
}
