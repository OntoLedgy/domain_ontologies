package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Artist struct {
	ID             int            `db:"id"`
	GID            sql.NullString `db:"gid"`
	Name           string         `db:"name"`
	SortName       string         `db:"sort_name"`
	BeginDateYear  sql.NullInt64  `db:"begin_date_year"`
	BeginDateMonth sql.NullInt64  `db:"begin_date_month"`
	BeginDateDay   sql.NullInt64  `db:"begin_date_day"`
	EndDateYear    sql.NullInt64  `db:"end_date_year"`
	EndDateMonth   sql.NullInt64  `db:"end_date_month"`
	EndDateDay     sql.NullInt64  `db:"end_date_day"`
	Type           sql.NullInt64  `db:"type"`
	Area           sql.NullInt64  `db:"area"`
	Gender         sql.NullInt64  `db:"gender"`
	Comment        string         `db:"comment"`
	EditsPending   int            `db:"edits_pending"`
	LastUpdated    pg.NullTime    `db:"last_updated"`
	Ended          bool           `db:"ended"`
	BeginArea      sql.NullInt64  `db:"begin_area"`
	EndArea        sql.NullInt64  `db:"end_area"`
}

func (Artist) GetTransactionObject() {

}
