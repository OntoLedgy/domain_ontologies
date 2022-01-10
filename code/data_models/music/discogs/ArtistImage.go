package dto

import (
	"database/sql"
)

type ArtistImage struct {
	ArtistID int            `db:"artist_id"`
	Type     sql.NullString `db:"type"`
	WIDth    sql.NullInt64  `db:"width"`
	Height   sql.NullInt64  `db:"height"`
}
