package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type LinkAttributeType struct {
	ID          int            `db:"id"`
	Parent      sql.NullInt64  `db:"parent"`
	Root        int            `db:"root"`
	ChildOrder  int            `db:"child_order"`
	GID         sql.NullString `db:"gid"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	LastUpdated pg.NullTime    `db:"last_updated"`
}
