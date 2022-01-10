package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionSeries struct {
	Collection int         `db:"collection"`
	Series     int         `db:"series"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
