package sparx_ea

import (
	"database/sql"
)

type TOperationpres struct {
	OperationID  int            `db:"operationid"`
	Precondition string         `db:"precondition"`
	Type         sql.NullString `db:"type"`
	Notes        sql.NullString `db:"notes"`
}
