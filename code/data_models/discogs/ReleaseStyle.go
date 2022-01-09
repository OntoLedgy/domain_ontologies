package dto

import (
	"database/sql"
)

type ReleaseStyle struct {
	ReleaseID int            `db:"release_id"`
	Style     sql.NullString `db:"style"`
}
