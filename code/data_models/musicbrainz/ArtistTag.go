package dto

import (
	pg "github.com/lib/pq"
)

type ArtistTag struct {
	Artist      int         `db:"artist"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
