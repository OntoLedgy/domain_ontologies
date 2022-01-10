package dto

import (
	"database/sql"
)

type TrackRaw struct {
	ID       int            `db:"id"`
	Release  int            `db:"release"`
	Title    string         `db:"title"`
	Artist   sql.NullString `db:"artist"`
	Sequence int            `db:"sequence"`
}
