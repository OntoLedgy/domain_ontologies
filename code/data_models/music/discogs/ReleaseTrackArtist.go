package dto

import (
	"database/sql"
)

type ReleaseTrackArtist struct {
	ID            int            `db:"id"`
	TrackID       sql.NullString `db:"track_id"`
	ReleaseID     int            `db:"release_id"`
	TrackSequence int            `db:"track_sequence"`
	ArtistID      int            `db:"artist_id"`
	ArtistName    sql.NullString `db:"artist_name"`
	Extra         bool           `db:"extra"`
	Anv           sql.NullString `db:"anv"`
	Position      sql.NullInt64  `db:"position"`
	JoinString    sql.NullString `db:"join_string"`
	Role          sql.NullString `db:"role"`
	Tracks        sql.NullString `db:"tracks"`
}
