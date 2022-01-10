package sparx_ea

import (
	"database/sql"
)

type TOperationposts struct {
	OperationID   int            `db:"operationid"`
	Postcondition string         `db:"postcondition"`
	Type          sql.NullString `db:"type"`
	Notes         sql.NullString `db:"notes"`
}
