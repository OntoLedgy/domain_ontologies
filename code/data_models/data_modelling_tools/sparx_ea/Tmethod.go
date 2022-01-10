package sparx_ea

import (
	"database/sql"
)

type TMethod struct {
	ObjectID int            `db:"object_id"`
	Name     string         `db:"name"`
	Scope    sql.NullString `db:"scope"`
	Type     sql.NullString `db:"type"`
}
