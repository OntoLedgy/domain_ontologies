package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type ReleaseMeta struct {
	ID               int            `db:"id"`
	DateAdded        pg.NullTime    `db:"date_added"`
	InfoURL          sql.NullString `db:"info_url"`
	AmazonAsin       sql.NullString `db:"amazon_asin"`
	AmazonStore      sql.NullString `db:"amazon_store"`
	CoverArtPresence sql.NullString `db:"cover_art_presence"`
}
