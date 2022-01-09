package dto

import (
	"database/sql"
)

type ReleaseImage struct {
	ReleaseID int            `db:"release_id"`
	Type      sql.NullString `db:"type"`
	WIDth     sql.NullInt64  `db:"width"`
	Height    sql.NullInt64  `db:"height"`
}
