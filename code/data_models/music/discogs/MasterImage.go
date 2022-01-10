package dto

import (
	"database/sql"
)

type MasterImage struct {
	MasterID int            `db:"master_id"`
	Type     sql.NullString `db:"type"`
	WIDth    sql.NullInt64  `db:"width"`
	Height   sql.NullInt64  `db:"height"`
}
