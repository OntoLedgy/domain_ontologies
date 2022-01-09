package dto

import (
	"database/sql"
)

type Master struct {
	ID          int            `db:"id"`
	Title       string         `db:"title"`
	Year        sql.NullInt64  `db:"year"`
	MainRelease int            `db:"main_release"`
	DataQuality sql.NullString `db:"data_quality"`
}
