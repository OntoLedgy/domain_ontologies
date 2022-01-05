package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionArea struct {
	Collection int         `db:"collection"`
	Area       int         `db:"area"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
