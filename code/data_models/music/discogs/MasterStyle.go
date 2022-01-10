package dto

import (
	"database/sql"
)

type MasterStyle struct {
	ID       int            `db:"id"`
	MasterID int            `db:"master_id"`
	Style    sql.NullString `db:"style"`
}
