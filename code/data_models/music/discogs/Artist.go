package dto

import (
	"database/sql"
)

type Artist struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Realname    sql.NullString `db:"realname"`
	Profile     sql.NullString `db:"profile"`
	DataQuality sql.NullString `db:"data_quality"`
}
