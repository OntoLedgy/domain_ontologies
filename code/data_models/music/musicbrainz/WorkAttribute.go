package dto

import (
	"database/sql"
)

type WorkAttribute struct {
	ID                            int            `db:"id"`
	Work                          int            `db:"work"`
	WorkAttributeType             int            `db:"work_attribute_type"`
	WorkAttributeTypeAllowedValue sql.NullInt64  `db:"work_attribute_type_allowed_value"`
	WorkAttributeText             sql.NullString `db:"work_attribute_text"`
}
