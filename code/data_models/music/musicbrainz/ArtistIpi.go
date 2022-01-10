package dto

import (
	pg "github.com/lib/pq"
)

type ArtistIpi struct {
	Artist       int         `db:"artist"`
	Ipi          string      `db:"ipi"`
	EditsPending int         `db:"edits_pending"`
	Created      pg.NullTime `db:"created"`
}
