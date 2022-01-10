package dto

import (
	"database/sql"
)

type AlternativeTrack struct {
	ID           int            `db:"id"`
	Name         sql.NullString `db:"name"`
	ArtistCredit sql.NullInt64  `db:"artist_credit"`
	RefCount     int            `db:"ref_count"`
}
