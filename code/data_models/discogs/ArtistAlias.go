package dto

import (
	"database/sql"
)

type ArtistAlias struct {
	ArtistID      int           `db:"artist_id"`
	AliasName     string        `db:"alias_name"`
	AliasArtistID sql.NullInt64 `db:"alias_artist_id"`
}
