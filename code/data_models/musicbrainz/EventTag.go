package dto

import (
	pg "github.com/lib/pq"
)

type EventTag struct {
	Event       int         `db:"event"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
