package dto

import "database/sql"

type CdtocRaw struct {
	ID            int            `db:"id"`
	Release       int            `db:"release"`
	DiscID        string         `db:"discid"`
	TrackCount    int            `db:"track_count"`
	LeadoutOffset int            `db:"leadout_offset"`
	TrackOffset   sql.NullString `db:"track_offset"`
}
