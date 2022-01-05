package dto

import (
	"database/sql"
)

type WorkMeta struct {
	ID          int           `db:"id"`
	Rating      sql.NullInt64 `db:"rating"`
	RatingCount sql.NullInt64 `db:"rating_count"`
}
