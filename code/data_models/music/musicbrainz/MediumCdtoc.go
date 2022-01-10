package dto

import (
	pg "github.com/lib/pq"
)

type MediumCdtoc struct {
	ID           int         `db:"id"`
	Medium       int         `db:"medium"`
	Cdtoc        int         `db:"cdtoc"`
	EditsPending int         `db:"edits_pending"`
	LastUpdated  pg.NullTime `db:"last_updated"`
}
