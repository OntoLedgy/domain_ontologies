package sparx_ea

import (
	"database/sql"
)

type UsysSystem struct {
	Property string         `db:"property"`
	Value    sql.NullString `db:"value"`
}
