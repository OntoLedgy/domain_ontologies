package dto

import (
	"database/sql"
)

type Language struct {
	ID        int            `db:"id"`
	IsoCode2t sql.NullString `db:"iso_code_2t"`
	IsoCode2b sql.NullString `db:"iso_code_2b"`
	IsoCode1  sql.NullString `db:"iso_code_1"`
	Name      string         `db:"name"`
	Frequency int            `db:"frequency"`
	IsoCode3  sql.NullString `db:"iso_code_3"`
}
