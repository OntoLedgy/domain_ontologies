package dto

import "database/sql"

type MediumIndex struct {
	Medium int            `db:"medium"`
	Toc    sql.NullString `db:"toc"`
}
