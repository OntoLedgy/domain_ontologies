package sparx_ea

import (
	"database/sql"
)

type TObjecttrx struct {
	ObjectID int             `db:"object_id"`
	Trx      string          `db:"trx"`
	Trxtype  string          `db:"trxtype"`
	Weight   sql.NullFloat64 `db:"weight"`
	Notes    sql.NullString  `db:"notes"`
}
