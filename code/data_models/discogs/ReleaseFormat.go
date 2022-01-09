package dto

import (
	"database/sql"
)

type ReleaseFormat struct {
	ID           int             `db:"id"`
	ReleaseID    int             `db:"release_id"`
	Name         sql.NullString  `db:"name"`
	Qty          sql.NullFloat64 `db:"qty"`
	TextString   sql.NullString  `db:"text_string"`
	Descriptions sql.NullString  `db:"descriptions"`
}
