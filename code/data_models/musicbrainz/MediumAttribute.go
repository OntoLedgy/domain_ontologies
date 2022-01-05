package dto

import (
	"database/sql"
)

type MediumAttribute struct {
	ID                              int            `db:"id"`
	Medium                          int            `db:"medium"`
	MediumAttributeType             int            `db:"medium_attribute_type"`
	MediumAttributeTypeAllowedValue sql.NullInt64  `db:"medium_attribute_type_allowed_value"`
	MediumAttributeText             sql.NullString `db:"medium_attribute_text"`
}
