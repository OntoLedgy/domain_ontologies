package sparx_ea

import (
	"database/sql"
)

type TDiagramobjects struct {
	DiagramID   sql.NullInt64  `db:"diagram_id"`
	ObjectID    sql.NullInt64  `db:"object_id"`
	Recttop     sql.NullInt64  `db:"recttop"`
	Rectleft    sql.NullInt64  `db:"rectleft"`
	Rectright   sql.NullInt64  `db:"rectright"`
	Rectbottom  sql.NullInt64  `db:"rectbottom"`
	Sequence    sql.NullInt64  `db:"sequence"`
	Objectstyle sql.NullString `db:"objectstyle"`
	InstanceID  int            `db:"instance_id"`
}
