package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type GenreAlias struct {
	ID               int            `db:"id"`
	Genre            int            `db:"genre"`
	Name             string         `db:"name"`
	Locale           sql.NullString `db:"locale"`
	EditsPending     int            `db:"edits_pending"`
	LastUpdated      pg.NullTime    `db:"last_updated"`
	PrimaryForLocale bool           `db:"primary_for_locale"`
}
