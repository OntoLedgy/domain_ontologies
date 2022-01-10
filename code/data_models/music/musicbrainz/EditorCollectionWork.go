package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionWork struct {
	Collection int         `db:"collection"`
	Work       int         `db:"work"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
