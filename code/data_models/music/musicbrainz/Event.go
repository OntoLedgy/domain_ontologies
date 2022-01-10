package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Event struct {
	ID             int            `db:"id"`
	GID            sql.NullString `db:"gid"`
	Name           string         `db:"name"`
	BeginDateYear  sql.NullInt64  `db:"begin_date_year"`
	BeginDateMonth sql.NullInt64  `db:"begin_date_month"`
	BeginDateDay   sql.NullInt64  `db:"begin_date_day"`
	EndDateYear    sql.NullInt64  `db:"end_date_year"`
	EndDateMonth   sql.NullInt64  `db:"end_date_month"`
	EndDateDay     sql.NullInt64  `db:"end_date_day"`
	Time           pg.NullTime    `db:"time"`
	Type           sql.NullInt64  `db:"type"`
	Cancelled      bool           `db:"cancelled"`
	Setlist        sql.NullString `db:"setlist"`
	Comment        string         `db:"comment"`
	EditsPending   int            `db:"edits_pending"`
	LastUpdated    pg.NullTime    `db:"last_updated"`
	Ended          bool           `db:"ended"`
}
