package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Label struct {
	ID             int            `db:"id"`
	GID            sql.NullString `db:"gid"`
	Name           string         `db:"name"`
	BeginDateYear  sql.NullInt64  `db:"begin_date_year"`
	BeginDateMonth sql.NullInt64  `db:"begin_date_month"`
	BeginDateDay   sql.NullInt64  `db:"begin_date_day"`
	EndDateYear    sql.NullInt64  `db:"end_date_year"`
	EndDateMonth   sql.NullInt64  `db:"end_date_month"`
	EndDateDay     sql.NullInt64  `db:"end_date_day"`
	LabelCode      sql.NullInt64  `db:"label_code"`
	Type           sql.NullInt64  `db:"type"`
	Area           sql.NullInt64  `db:"area"`
	Comment        string         `db:"comment"`
	EditsPending   int            `db:"edits_pending"`
	LastUpdated    pg.NullTime    `db:"last_updated"`
	Ended          bool           `db:"ended"`
}
