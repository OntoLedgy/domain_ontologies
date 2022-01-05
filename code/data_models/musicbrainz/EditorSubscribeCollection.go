package dto

import (
	"database/sql"
)

type EditorSubscribeCollection struct {
	ID           int            `db:"id"`
	Editor       int            `db:"editor"`
	Collection   int            `db:"collection"`
	LastEditSent int            `db:"last_edit_sent"`
	Available    bool           `db:"available"`
	LastSeenName sql.NullString `db:"last_seen_name"`
}
