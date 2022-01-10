package dto

import (
	pg "github.com/lib/pq"
)

type PlaceTag struct {
	Place       int         `db:"place"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
