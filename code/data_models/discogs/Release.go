package dto

import (
	"database/sql"
)

type Release struct {
	ID          int            `db:"id"`
	Title       string         `db:"title"`
	Released    sql.NullString `db:"released"`
	Country     sql.NullString `db:"country"`
	Notes       sql.NullString `db:"notes"`
	DataQuality sql.NullString `db:"data_quality"`
	Main        sql.NullInt64  `db:"main"`
	MasterID    sql.NullInt64  `db:"master_id"`
	Status      sql.NullString `db:"status"`
}
