package dto

import (
	"database/sql"
	"time"
)

type DeletedEntity struct {
	GID       sql.NullString `db:"gid"`
	Data      sql.NullString `db:"data"`
	DeletedAt time.Time      `db:"deleted_at"`
}
