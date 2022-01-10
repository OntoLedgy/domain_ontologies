package dto

import (
	"database/sql"
)

type RecordingAttribute struct {
	ID                                 int            `db:"id"`
	Recording                          int            `db:"recording"`
	RecordingAttributeType             int            `db:"recording_attribute_type"`
	RecordingAttributeTypeAllowedValue sql.NullInt64  `db:"recording_attribute_type_allowed_value"`
	RecordingAttributeText             sql.NullString `db:"recording_attribute_text"`
}
