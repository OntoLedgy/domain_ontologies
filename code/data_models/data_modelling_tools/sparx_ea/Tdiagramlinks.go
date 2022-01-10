package sparx_ea

import (
	"database/sql"
)

type TDiagramlinks struct {
	DiagramID   sql.NullInt64  `db:"diagramid"`
	ConnectorID sql.NullInt64  `db:"connectorid"`
	Geometry    sql.NullString `db:"geometry"`
	Style       sql.NullString `db:"style"`
	HIDden      int            `db:"hidden"`
	Path        sql.NullString `db:"path"`
	InstanceID  int            `db:"instance_id"`
}
