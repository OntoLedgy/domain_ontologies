package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Medium struct {
	ID           int           `db:"id"`
	Release      int           `db:"release"`
	Position     int           `db:"position"`
	Format       sql.NullInt64 `db:"format"`
	Name         string        `db:"name"`
	EditsPending int           `db:"edits_pending"`
	LastUpdated  pg.NullTime   `db:"last_updated"`
	TrackCount   int           `db:"track_count"`
}
