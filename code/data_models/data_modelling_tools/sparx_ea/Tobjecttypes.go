package sparx_ea

import (
	"database/sql"
)

type TObjecttypes struct {
	ObjectType   string         `db:"object_type"`
	Description  sql.NullString `db:"description"`
	Designobject int            `db:"designobject"`
	ImageID      sql.NullInt64  `db:"imageid"`
}
