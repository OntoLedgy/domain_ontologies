package dto

import (
	pg "github.com/lib/pq"
)

type EditorCollectionDeletedEntity struct {
	Collection int            `db:"collection"`
	GID        sql.NullString `db:"gid"`
	Added      pg.NullTime    `db:"added"`
	Position   int            `db:"position"`
	Comment    string         `db:"comment"`
}
