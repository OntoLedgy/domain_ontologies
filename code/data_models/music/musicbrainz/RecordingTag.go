package dto

import (
	pg "github.com/lib/pq"
)

type RecordingTag struct {
	Recording   int         `db:"recording"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
