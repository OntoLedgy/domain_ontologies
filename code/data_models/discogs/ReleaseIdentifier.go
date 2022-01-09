package dto

import (
	"database/sql"
)

type ReleaseIdentifier struct {
	ID          int            `db:"id"`
	ReleaseID   int            `db:"release_id"`
	Description sql.NullString `db:"description"`
	Type        sql.NullString `db:"type"`
	Value       sql.NullString `db:"value"`
}
