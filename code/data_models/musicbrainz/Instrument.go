package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Instrument struct {
	ID           int            `db:"id"`
	GID          sql.NullString `db:"gid"`
	Name         string         `db:"name"`
	Type         sql.NullInt64  `db:"type"`
	EditsPending int            `db:"edits_pending"`
	LastUpdated  pg.NullTime    `db:"last_updated"`
	Comment      string         `db:"comment"`
	Description  string         `db:"description"`
}
