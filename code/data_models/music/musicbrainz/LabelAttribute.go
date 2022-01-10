package dto

import (
	"database/sql"
)

type LabelAttribute struct {
	ID                             int            `db:"id"`
	Label                          int            `db:"label"`
	LabelAttributeType             int            `db:"label_attribute_type"`
	LabelAttributeTypeAllowedValue sql.NullInt64  `db:"label_attribute_type_allowed_value"`
	LabelAttributeText             sql.NullString `db:"label_attribute_text"`
}
