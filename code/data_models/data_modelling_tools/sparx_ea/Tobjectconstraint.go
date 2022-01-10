package sparx_ea

import (
	"database/sql"
)

type TObjectconstraint struct {
	ObjectID       int             `db:"object_id"`
	Constraint     string          `db:"Constraint"`
	Constrainttype string          `db:"constrainttype"`
	Weight         sql.NullFloat64 `db:"weight"`
	Notes          sql.NullString  `db:"notes"`
	Status         sql.NullString  `db:"status"`
}
