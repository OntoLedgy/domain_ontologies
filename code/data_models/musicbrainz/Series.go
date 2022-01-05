package dto

import (
	"database/sql"
	pg "github.com/lib/pq"
)

type Series struct {
	ID                int            `db:"id"`
	GID               sql.NullString `db:"gid"`
	Name              string         `db:"name"`
	Comment           string         `db:"comment"`
	Type              int            `db:"type"`
	OrderingAttribute int            `db:"ordering_attribute"`
	OrderingType      int            `db:"ordering_type"`
	EditsPending      int            `db:"edits_pending"`
	LastUpdated       pg.NullTime    `db:"last_updated"`
}
