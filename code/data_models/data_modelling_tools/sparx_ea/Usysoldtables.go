package sparx_ea

import (
	"database/sql"
)

type Usysoldtables struct {
	Tablename sql.NullString `db:"tablename"`
	Newname   sql.NullString `db:"newname"`
	Relorder  sql.NullInt64  `db:"relorder"`
	Fixcode   int            `db:"fixcode"`
}
