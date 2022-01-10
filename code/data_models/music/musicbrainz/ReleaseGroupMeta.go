package dto

import (
	"database/sql"
)

type ReleaseGroupMeta struct {
	ID                    int           `db:"id"`
	ReleaseCount          int           `db:"release_count"`
	FirstReleaseDateYear  sql.NullInt64 `db:"first_release_date_year"`
	FirstReleaseDateMonth sql.NullInt64 `db:"first_release_date_month"`
	FirstReleaseDateDay   sql.NullInt64 `db:"first_release_date_day"`
	Rating                sql.NullInt64 `db:"rating"`
	RatingCount           sql.NullInt64 `db:"rating_count"`
}
