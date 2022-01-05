package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionArtist struct {
	Collection int         `db:"collection"`
	Artist     int         `db:"artist"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
