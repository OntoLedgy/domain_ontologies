package sparx_ea

import (
	"database/sql"
)

type Usystables struct {
	Tablename   string         `db:"tablename"`
	Relorder    sql.NullInt64  `db:"relorder"`
	Displayname sql.NullString `db:"displayname"`
	Fromver     sql.NullString `db:"fromver"`
	Tover       sql.NullString `db:"tover"`
}
