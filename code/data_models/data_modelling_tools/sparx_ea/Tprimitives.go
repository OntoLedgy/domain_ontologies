package sparx_ea

import (
	"database/sql"
)

type TPrimitives struct {
	Datatype    string         `db:"datatype"`
	Description sql.NullString `db:"description"`
}
