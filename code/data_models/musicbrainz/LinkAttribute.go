package dto

import (
	pg "github.com/lib/pq"
)

type LinkAttribute struct {
	Link          int         `db:"link"`
	AttributeType int         `db:"attribute_type"`
	Created       pg.NullTime `db:"created"`
}
