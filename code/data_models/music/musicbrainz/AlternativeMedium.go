package dto

import (
	"database/sql"
)

type AlternativeMedium struct {
	ID                 int            `db:"id"`
	Medium             int            `db:"medium"`
	AlternativeRelease int            `db:"alternative_release"`
	Name               sql.NullString `db:"name"`
}
