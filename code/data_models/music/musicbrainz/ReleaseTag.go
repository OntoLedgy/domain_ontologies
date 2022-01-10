package dto

import (
	pg "github.com/lib/pq"
)

type ReleaseTag struct {
	Release     int         `db:"release"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
