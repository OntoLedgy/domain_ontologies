package dto

import (
	"database/sql"
	"time"

	pg "github.com/lib/pq"
)

type Edit struct {
	ID         int           `db:"id"`
	Editor     int           `db:"editor"`
	Type       int           `db:"type"`
	Status     int           `db:"status"`
	Autoedit   int           `db:"autoedit"`
	OpenTime   pg.NullTime   `db:"open_time"`
	CloseTime  pg.NullTime   `db:"close_time"`
	ExpireTime time.Time     `db:"expire_time"`
	Language   sql.NullInt64 `db:"language"`
	Quality    int           `db:"quality"`
}
