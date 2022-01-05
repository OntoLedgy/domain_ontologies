package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionInstrument struct {
	Collection int         `db:"collection"`
	Instrument int         `db:"instrument"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
