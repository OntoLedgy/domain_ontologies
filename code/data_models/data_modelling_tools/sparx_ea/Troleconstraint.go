package sparx_ea

import (
	"database/sql"
)

type TRoleconstraint struct {
	ConnectorID    int            `db:"connectorid"`
	Constraint     string         `db:"Constraint"`
	Connectorend   string         `db:"connectorend"`
	Constrainttype string         `db:"constrainttype"`
	Notes          sql.NullString `db:"notes"`
}
