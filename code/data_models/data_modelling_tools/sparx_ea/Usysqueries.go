package sparx_ea

import (
	"database/sql"
)

type Usysqueries struct {
	Queryname sql.NullString `db:"queryname"`
	Newname   sql.NullString `db:"newname"`
	Fixcode   int            `db:"fixcode"`
}
