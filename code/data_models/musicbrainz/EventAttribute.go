package dto

import (
	"database/sql"
)

type EventAttribute struct {
	ID                             int            `db:"id"`
	Event                          int            `db:"event"`
	EventAttributeType             int            `db:"event_attribute_type"`
	EventAttributeTypeAllowedValue sql.NullInt64  `db:"event_attribute_type_allowed_value"`
	EventAttributeText             sql.NullString `db:"event_attribute_text"`
}
