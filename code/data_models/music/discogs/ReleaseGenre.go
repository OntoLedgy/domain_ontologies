package dto

import (
	"database/sql"
)

type ReleaseGenre struct {
	ID        int            `db:"id"`
	ReleaseID int            `db:"release_id"`
	Genre     sql.NullString `db:"genre"`
}
