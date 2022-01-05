package dto

import (
	"database/sql"
)

type ArtistAttribute struct {
	ID                              int            `db:"id"`
	Artist                          int            `db:"artist"`
	ArtistAttributeType             int            `db:"artist_attribute_type"`
	ArtistAttributeTypeAllowedValue sql.NullInt64  `db:"artist_attribute_type_allowed_value"`
	ArtistAttributeText             sql.NullString `db:"artist_attribute_text"`
}
