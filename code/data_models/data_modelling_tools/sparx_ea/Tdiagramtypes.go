package sparx_ea

import (
	"database/sql"
)

type TDiagramtypes struct {
	DiagramType string         `db:"diagram_type"`
	Name        sql.NullString `db:"name"`
	PackageID   sql.NullInt64  `db:"package_id"`
}
