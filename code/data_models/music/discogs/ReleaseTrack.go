package dto

import (
	"database/sql"
)

type ReleaseTrack struct {
	ID        int            `db:"id"`
	ReleaseID int            `db:"release_id"`
	Sequence  int            `db:"sequence"`
	Position  sql.NullString `db:"position"`
	Parent    sql.NullInt64  `db:"parent"`
	Title     sql.NullString `db:"title"`
	Duration  sql.NullString `db:"duration"`
	TrackID   sql.NullString `db:"track_id"`
}
