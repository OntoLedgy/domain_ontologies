package sparx_ea

import (
	"database/sql"
)

type TStatustypes struct {
	Status      string         `db:"status"`
	Description sql.NullString `db:"description"`
}
