package dto

import (
	"database/sql"
)

type ReleaseCompany struct {
	ID             int            `db:"id"`
	ReleaseID      int            `db:"release_id"`
	CompanyID      int            `db:"company_id"`
	CompanyName    string         `db:"company_name"`
	EntityType     sql.NullString `db:"entity_type"`
	EntityTypeName sql.NullString `db:"entity_type_name"`
	Uri            sql.NullString `db:"uri"`
}
