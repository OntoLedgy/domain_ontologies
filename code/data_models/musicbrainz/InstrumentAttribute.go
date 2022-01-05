package dto

import (
	"database/sql"
)

type InstrumentAttribute struct {
	ID                                  int            `db:"id"`
	Instrument                          int            `db:"instrument"`
	InstrumentAttributeType             int            `db:"instrument_attribute_type"`
	InstrumentAttributeTypeAllowedValue sql.NullInt64  `db:"instrument_attribute_type_allowed_value"`
	InstrumentAttributeText             sql.NullString `db:"instrument_attribute_text"`
}
