package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionRelease struct {
	Collection int         `db:"collection"`
	Release    int         `db:"release"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
