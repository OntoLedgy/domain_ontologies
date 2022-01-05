package dto

import (
	"database/sql"
)

type ReleaseCountry struct {
	Release   int           `db:"release"`
	Country   int           `db:"country"`
	DateYear  sql.NullInt64 `db:"date_year"`
	DateMonth sql.NullInt64 `db:"date_month"`
	DateDay   sql.NullInt64 `db:"date_day"`
}
