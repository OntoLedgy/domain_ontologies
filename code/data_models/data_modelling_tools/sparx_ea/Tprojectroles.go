package sparx_ea

import (
	"database/sql"
)

type TProjectroles struct {
	Role        string         `db:"role"`
	Description sql.NullString `db:"description"`
	Notes       sql.NullString `db:"notes"`
}
