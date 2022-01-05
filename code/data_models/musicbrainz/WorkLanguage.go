package dto

import (
	pg "github.com/lib/pq"
)

type WorkLanguage struct {
	Work         int         `db:"work"`
	Language     int         `db:"language"`
	EditsPending int         `db:"edits_pending"`
	Created      pg.NullTime `db:"created"`
}
