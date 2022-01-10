package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionRecording struct {
	Collection int         `db:"collection"`
	Recording  int         `db:"recording"`
	Added      pg.NullTime `db:"added"`
	Position   int         `db:"position"`
	Comment    string      `db:"comment"`
}
