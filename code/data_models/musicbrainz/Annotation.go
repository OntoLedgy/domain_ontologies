package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Annotation struct {
	ID        int            `db:"id"`
	Editor    int            `db:"editor"`
	Text      sql.NullString `db:"text"`
	Changelog sql.NullString `db:"changelog"`
	Created   pg.NullTime    `db:"created"`
}
