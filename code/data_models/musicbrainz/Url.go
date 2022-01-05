package dto

import (
	"database/sql"
	pg "github.com/lib/pq"
)

type Url struct {
	ID           int            `db:"id"`
	GID          sql.NullString `db:"gid"`
	URL          string         `db:"url"`
	EditsPending int            `db:"edits_pending"`
	LastUpdated  pg.NullTime    `db:"last_updated"`
}
