package dto

import (
	pg "github.com/lib/pq"
)

type ArtistIsni struct {
	Artist       int         `db:"artist"`
	Isni         string      `db:"isni"`
	EditsPending int         `db:"edits_pending"`
	Created      pg.NullTime `db:"created"`
}
