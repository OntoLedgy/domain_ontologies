package sparx_ea

import (
	"database/sql"
)

type TConnectorconstraint struct {
	ConnectorID    int            `db:"connectorid"`
	Constraint     string         `db:"Constraint"`
	Constrainttype sql.NullString `db:"constrainttype"`
	Notes          sql.NullString `db:"notes"`
}
