package sparx_ea

import (
	"database/sql"
)

type TImplement struct {
	Type sql.NullString `db:"type"`
}
