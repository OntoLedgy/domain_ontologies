package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionPlace struct {
	Collection int         `db:"collection"`
	Place      int         `db:"place"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
