package dto

import (
	pg "github.com/lib/pq"
)

type WorkTag struct {
	Work        int         `db:"work"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
