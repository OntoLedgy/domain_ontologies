package dto

import (
	pg "github.com/lib/pq"
)

type TagRelation struct {
	Tag1        int         `db:"tag1"`
	Tag2        int         `db:"tag2"`
	Weight      int         `db:"weight"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
