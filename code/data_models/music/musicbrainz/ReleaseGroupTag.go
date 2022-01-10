package dto

import (
	pg "github.com/lib/pq"
)

type ReleaseGroupTag struct {
	ReleaseGroup int         `db:"release_group"`
	Tag          int         `db:"tag"`
	Count        int         `db:"count"`
	LastUpdated  pg.NullTime `db:"last_updated"`
}
