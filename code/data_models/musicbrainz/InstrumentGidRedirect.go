package dto

import (
	"database/sql"
	pg "github.com/lib/pq"
)

type InstrumentGidRedirect struct {
	GID     sql.NullString `db:"gid"`
	NewID   int            `db:"new_id"`
	Created pg.NullTime    `db:"created"`
}
