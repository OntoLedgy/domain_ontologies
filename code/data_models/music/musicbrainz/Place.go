package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Place struct {
	ID             int            `db:"id"`
	GID            sql.NullString `db:"gid"`
	Name           string         `db:"name"`
	Type           sql.NullInt64  `db:"type"`
	Address        string         `db:"address"`
	Area           sql.NullInt64  `db:"area"`
	Coordinates    sql.NullString `db:"coordinates"`
	Comment        string         `db:"comment"`
	EditsPending   int            `db:"edits_pending"`
	LastUpdated    pg.NullTime    `db:"last_updated"`
	BeginDateYear  sql.NullInt64  `db:"begin_date_year"`
	BeginDateMonth sql.NullInt64  `db:"begin_date_month"`
	BeginDateDay   sql.NullInt64  `db:"begin_date_day"`
	EndDateYear    sql.NullInt64  `db:"end_date_year"`
	EndDateMonth   sql.NullInt64  `db:"end_date_month"`
	EndDateDay     sql.NullInt64  `db:"end_date_day"`
	Ended          bool           `db:"ended"`
}
