package dto

import (
	pg "github.com/lib/pq"
)

type LabelIpi struct {
	Label        int         `db:"label"`
	Ipi          string      `db:"ipi"`
	EditsPending int         `db:"edits_pending"`
	Created      pg.NullTime `db:"created"`
}
