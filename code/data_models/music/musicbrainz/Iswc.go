package dto

import (
	"database/sql"
	"time"
)

type Iswc struct {
	ID           int            `db:"id"`
	Work         int            `db:"work"`
	Iswc         sql.NullString `db:"iswc"`
	Source       sql.NullInt64  `db:"source"`
	EditsPending int            `db:"edits_pending"`
	Created      time.Time      `db:"created"`
}
