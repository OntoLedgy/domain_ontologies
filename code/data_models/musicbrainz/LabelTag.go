package dto

import (
	pg "github.com/lib/pq"
)

type LabelTag struct {
	Label       int         `db:"label"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
