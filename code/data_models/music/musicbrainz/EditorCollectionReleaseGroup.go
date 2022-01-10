package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionReleaseGroup struct {
	Collection   int         `db:"collection"`
	ReleaseGroup int         `db:"release_group"`
	Added        pg.NullTime `db:"added"`
	Position     int         `db:"position"`
	Comment      string      `db:"comment"`
}
