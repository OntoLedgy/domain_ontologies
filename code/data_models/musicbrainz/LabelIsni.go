package dto

import (
	pg "github.com/lib/pq"
)

type LabelIsni struct {
	Label        int         `db:"label"`
	Isni         string      `db:"isni"`
	EditsPending int         `db:"edits_pending"`
	Created      pg.NullTime `db:"created"`
}
