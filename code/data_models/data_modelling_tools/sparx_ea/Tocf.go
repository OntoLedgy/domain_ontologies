package sparx_ea

import (
	"database/sql"
)

type TOcf struct {
	Objecttype       sql.NullString  `db:"objecttype"`
	Complexityweight sql.NullFloat64 `db:"complexityweight"`
}
