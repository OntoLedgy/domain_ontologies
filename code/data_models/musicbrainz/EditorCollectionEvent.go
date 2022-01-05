package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionEvent struct {
	Collection int         `db:"collection"`
	Event      int         `db:"event"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
