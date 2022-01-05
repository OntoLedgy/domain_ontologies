package dto

import (
	"database/sql"
)

type ReleaseGroupAttribute struct {
	ID                                    int            `db:"id"`
	ReleaseGroup                          int            `db:"release_group"`
	ReleaseGroupAttributeType             int            `db:"release_group_attribute_type"`
	ReleaseGroupAttributeTypeAllowedValue sql.NullInt64  `db:"release_group_attribute_type_allowed_value"`
	ReleaseGroupAttributeText             sql.NullString `db:"release_group_attribute_text"`
}
