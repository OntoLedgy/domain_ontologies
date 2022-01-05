package dto

import (
	"database/sql"
	pg "github.com/lib/pq"
)

type Cdtoc struct {
	ID            int            `db:"id"`
	DiscID        string         `db:"discid"`
	FreedbID      string         `db:"freedb_id"`
	TrackCount    int            `db:"track_count"`
	LeadoutOffset int            `db:"leadout_offset"`
	TrackOffset   sql.NullString `db:"track_offset"`
	Degraded      bool           `db:"degraded"`
	Created       pg.NullTime    `db:"created"`
}
