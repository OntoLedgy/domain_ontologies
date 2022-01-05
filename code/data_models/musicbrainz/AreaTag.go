package dto

import (
	pg "github.com/lib/pq"
)

type AreaTag struct {
	Area        int         `db:"area"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
