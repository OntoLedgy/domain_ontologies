package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionLabel struct {
	Collection int         `db:"collection"`
	Label      int         `db:"label"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
