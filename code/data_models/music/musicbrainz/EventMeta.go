package dto

import (
	"database/sql"
)

type EventMeta struct {
	ID               int            `db:"id"`
	Rating           sql.NullInt64  `db:"rating"`
	RatingCount      sql.NullInt64  `db:"rating_count"`
	EventArtPresence sql.NullString `db:"event_art_presence"`
}
