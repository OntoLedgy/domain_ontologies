package sparx_ea

import (
	"database/sql"
	"time"
)

type TSeclocks struct {
	UserID     string         `db:"userid"`
	GroupID    sql.NullString `db:"groupid"`
	Entitytype string         `db:"entitytype"`
	EntityID   string         `db:"entityid"`
	Timestamp  time.Time      `db:"timestamp"`
	Locktype   sql.NullString `db:"locktype"`
}
