package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type ReleaseCoverart struct {
	ID          int            `db:"id"`
	LastUpdated pg.NullTime    `db:"last_updated"`
	CoverArtURL sql.NullString `db:"cover_art_url"`
}
