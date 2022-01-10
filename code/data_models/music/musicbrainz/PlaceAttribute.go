package dto

import (
	"database/sql"
)

type PlaceAttribute struct {
	ID                             int            `db:"id"`
	Place                          int            `db:"place"`
	PlaceAttributeType             int            `db:"place_attribute_type"`
	PlaceAttributeTypeAllowedValue sql.NullInt64  `db:"place_attribute_type_allowed_value"`
	PlaceAttributeText             sql.NullString `db:"place_attribute_text"`
}
