package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type LinkTypeAttributeType struct {
	LinkType      int           `db:"link_type"`
	AttributeType int           `db:"attribute_type"`
	Min           sql.NullInt64 `db:"min"`
	Max           sql.NullInt64 `db:"max"`
	LastUpdated   pg.NullTime   `db:"last_updated"`
}
