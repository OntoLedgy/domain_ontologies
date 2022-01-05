package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type ReleaseLabel struct {
	ID            int            `db:"id"`
	Release       int            `db:"release"`
	Label         sql.NullInt64  `db:"label"`
	CatalogNumber sql.NullString `db:"catalog_number"`
	LastUpdated   pg.NullTime    `db:"last_updated"`
}
