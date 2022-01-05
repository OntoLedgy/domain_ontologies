package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Track struct {
	ID           int            `db:"id"`
	GID          sql.NullString `db:"gid"`
	Recording    int            `db:"recording"`
	Medium       int            `db:"medium"`
	Position     int            `db:"position"`
	Number       string         `db:"number"`
	Name         string         `db:"name"`
	ArtistCredit int            `db:"artist_credit"`
	Length       sql.NullInt64  `db:"length"`
	EditsPending int            `db:"edits_pending"`
	LastUpdated  pg.NullTime    `db:"last_updated"`
	IsDataTrack  bool           `db:"is_data_track"`
}
