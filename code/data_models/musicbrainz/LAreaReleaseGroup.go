package dto

import (
	pg "github.com/lib/pq"
)

type LAreaReleaseGroup struct {
	ID            int         `db:"id"`
	Link          int         `db:"link"`
	Entity0       int         `db:"entity0"`
	Entity1       int         `db:"entity1"`
	EditsPending  int         `db:"edits_pending"`
	LastUpdated   pg.NullTime `db:"last_updated"`
	LinkOrder     int         `db:"link_order"`
	Entity0Credit string      `db:"entity0_credit"`
	Entity1Credit string      `db:"entity1_credit"`
}
