package sparx_ea

import (
	"database/sql"
)

type TXrefuser struct {
	XrefID      string         `db:"xrefid"`
	Name        sql.NullString `db:"name"`
	Type        sql.NullString `db:"type"`
	Namespace   sql.NullString `db:"namespace"`
	Requirement sql.NullString `db:"requirement"`
	Constraint  sql.NullString `db:"Constraint"`
	Behavior    sql.NullString `db:"behavior"`
	Partition   sql.NullString `db:"partition"`
	Description sql.NullString `db:"description"`
	Client      sql.NullString `db:"client"`
	Supplier    sql.NullString `db:"supplier"`
	Link        sql.NullString `db:"link"`
	ToolID      sql.NullString `db:"toolid"`
	Visibility  sql.NullString `db:"visibility"`
}
