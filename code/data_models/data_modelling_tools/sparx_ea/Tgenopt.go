package sparx_ea

import (
	"database/sql"
)

type TGenopt struct {
	Appliesto sql.NullString `db:"appliesto"`
	Option    sql.NullString `db:"option"`
}
