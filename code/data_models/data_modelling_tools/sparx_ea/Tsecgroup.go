package sparx_ea

import (
	"database/sql"
)

type TSecgroup struct {
	GroupID     string         `db:"groupid"`
	Groupname   string         `db:"groupname"`
	Description sql.NullString `db:"description"`
}
