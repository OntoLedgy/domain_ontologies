package dto

import (
	"database/sql"
)

type MasterGenre struct {
	ID       int            `db:"id"`
	MasterID int            `db:"master_id"`
	Genre    sql.NullString `db:"genre"`
}
