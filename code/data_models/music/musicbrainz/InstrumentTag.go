package dto

import (
	pg "github.com/lib/pq"
)

type InstrumentTag struct {
	Instrument  int         `db:"instrument"`
	Tag         int         `db:"tag"`
	Count       int         `db:"count"`
	LastUpdated pg.NullTime `db:"last_updated"`
}
