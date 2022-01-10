package dto

import (
	"database/sql"
)

type MasterArtist struct {
	ID         int            `db:"id"`
	MasterID   int            `db:"master_id"`
	ArtistID   int            `db:"artist_id"`
	ArtistName sql.NullString `db:"artist_name"`
	Anv        sql.NullString `db:"anv"`
	Position   sql.NullInt64  `db:"position"`
	JoinString sql.NullString `db:"join_string"`
	Role       sql.NullString `db:"role"`
}
