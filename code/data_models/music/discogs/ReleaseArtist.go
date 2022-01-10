package dto

import (
	"database/sql"
)

type ReleaseArtist struct {
	ID         int            `db:"id"`
	ReleaseID  int            `db:"release_id"`
	ArtistID   int            `db:"artist_id"`
	ArtistName sql.NullString `db:"artist_name"`
	Extra      int            `db:"extra"`
	Anv        sql.NullString `db:"anv"`
	Position   sql.NullInt64  `db:"position"`
	JoinString sql.NullString `db:"join_string"`
	Role       sql.NullString `db:"role"`
	Tracks     sql.NullString `db:"tracks"`
}
