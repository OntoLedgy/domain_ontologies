package dto

import (
	"database/sql"
)

type ReleaseAttribute struct {
	ID                               int            `db:"id"`
	Release                          int            `db:"release"`
	ReleaseAttributeType             int            `db:"release_attribute_type"`
	ReleaseAttributeTypeAllowedValue sql.NullInt64  `db:"release_attribute_type_allowed_value"`
	ReleaseAttributeText             sql.NullString `db:"release_attribute_text"`
}
