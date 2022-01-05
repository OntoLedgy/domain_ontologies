package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type LinkType struct {
	ID                 int            `db:"id"`
	Parent             sql.NullInt64  `db:"parent"`
	ChildOrder         int            `db:"child_order"`
	GID                sql.NullString `db:"gid"`
	EntityType0        string         `db:"entity_type0"`
	EntityType1        string         `db:"entity_type1"`
	Name               string         `db:"name"`
	Description        sql.NullString `db:"description"`
	LinkPhrase         string         `db:"link_phrase"`
	ReverseLinkPhrase  string         `db:"reverse_link_phrase"`
	LongLinkPhrase     string         `db:"long_link_phrase"`
	Priority           int            `db:"priority"`
	LastUpdated        pg.NullTime    `db:"last_updated"`
	IsDeprecated       bool           `db:"is_deprecated"`
	HasDates           bool           `db:"has_dates"`
	Entity0Cardinality int            `db:"entity0_cardinality"`
	Entity1Cardinality int            `db:"entity1_cardinality"`
}
