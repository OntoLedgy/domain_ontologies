package dto

import (
	pg "github.com/lib/pq"
)

type SeriesTag struct {
	Series      int         `db:"series"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
