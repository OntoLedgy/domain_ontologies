package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Isrc struct {
	ID           int           `db:"id"`
	Recording    int           `db:"recording"`
	Isrc         string        `db:"isrc"`
	Source       sql.NullInt64 `db:"source"`
	EditsPending int           `db:"edits_pending"`
	Created      pg.NullTime   `db:"created"`
}
