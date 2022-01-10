package sparx_ea

import (
	"database/sql"
)

type TOperationparams struct {
	OperationID int            `db:"operationid"`
	Name        string         `db:"name"`
	Type        sql.NullString `db:"type"`
	Default     sql.NullString `db:"Default"`
	Notes       sql.NullString `db:"notes"`
	Pos         sql.NullInt64  `db:"pos"`
	Const       sql.NullInt64  `db:"const"`
	Style       sql.NullString `db:"style"`
	Kind        sql.NullString `db:"kind"`
	Classifier  sql.NullString `db:"classifier"`
	EaGuID      sql.NullString `db:"ea_guid"`
	Styleex     sql.NullString `db:"styleex"`
}
