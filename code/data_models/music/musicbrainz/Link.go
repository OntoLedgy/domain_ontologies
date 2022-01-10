package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Link struct {
	ID             int           `db:"id"`
	LinkType       int           `db:"link_type"`
	BeginDateYear  sql.NullInt64 `db:"begin_date_year"`
	BeginDateMonth sql.NullInt64 `db:"begin_date_month"`
	BeginDateDay   sql.NullInt64 `db:"begin_date_day"`
	EndDateYear    sql.NullInt64 `db:"end_date_year"`
	EndDateMonth   sql.NullInt64 `db:"end_date_month"`
	EndDateDay     sql.NullInt64 `db:"end_date_day"`
	AttributeCount int           `db:"attribute_count"`
	Created        pg.NullTime   `db:"created"`
	Ended          bool          `db:"ended"`
}
