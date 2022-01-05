package dto

import (
	"database/sql"
)

type SeriesAttribute struct {
	ID                              int            `db:"id"`
	Series                          int            `db:"series"`
	SeriesAttributeType             int            `db:"series_attribute_type"`
	SeriesAttributeTypeAllowedValue sql.NullInt64  `db:"series_attribute_type_allowed_value"`
	SeriesAttributeText             sql.NullString `db:"series_attribute_text"`
}
