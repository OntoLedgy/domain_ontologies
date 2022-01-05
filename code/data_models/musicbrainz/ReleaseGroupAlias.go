package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type ReleaseGroupAlias struct {
	ID               int            `db:"id"`
	ReleaseGroup     int            `db:"release_group"`
	Name             string         `db:"name"`
	Locale           sql.NullString `db:"locale"`
	EditsPending     int            `db:"edits_pending"`
	LastUpdated      pg.NullTime    `db:"last_updated"`
	Type             sql.NullInt64  `db:"type"`
	SortName         string         `db:"sort_name"`
	BeginDateYear    sql.NullInt64  `db:"begin_date_year"`
	BeginDateMonth   sql.NullInt64  `db:"begin_date_month"`
	BeginDateDay     sql.NullInt64  `db:"begin_date_day"`
	EndDateYear      sql.NullInt64  `db:"end_date_year"`
	EndDateMonth     sql.NullInt64  `db:"end_date_month"`
	EndDateDay       sql.NullInt64  `db:"end_date_day"`
	PrimaryForLocale bool           `db:"primary_for_locale"`
	Ended            bool           `db:"ended"`
}
