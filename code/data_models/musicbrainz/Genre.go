package dto

import (
	"database/sql"
	pg "github.com/lib/pq"
)

type Genre struct {
	ID           int            `db:"id"`
	GID          sql.NullString `db:"gid"`
	Name         string         `db:"name"`
	Comment      string         `db:"comment"`
	EditsPending int            `db:"edits_pending"`
	LastUpdated  pg.NullTime    `db:"last_updated"`
}
