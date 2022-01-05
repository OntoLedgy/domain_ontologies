package dto

import (
	"database/sql"
)

type ReleaseUnknownCountry struct {
	Release   int           `db:"release"`
	DateYear  sql.NullInt64 `db:"date_year"`
	DateMonth sql.NullInt64 `db:"date_month"`
	DateDay   sql.NullInt64 `db:"date_day"`
}
