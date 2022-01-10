package dto

import (
	"database/sql"
)

type AreaAttribute struct {
	ID                            int            `db:"id"`
	Area                          int            `db:"area"`
	AreaAttributeType             int            `db:"area_attribute_type"`
	AreaAttributeTypeAllowedValue sql.NullInt64  `db:"area_attribute_type_allowed_value"`
	AreaAttributeText             sql.NullString `db:"area_attribute_text"`
}
