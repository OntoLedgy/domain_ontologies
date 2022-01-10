package dto

import "database/sql"

type EditData struct {
	Edit int            `db:"edit"`
	Data sql.NullString `db:"data"`
}
