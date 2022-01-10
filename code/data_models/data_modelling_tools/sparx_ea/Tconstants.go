package sparx_ea

import (
	"database/sql"
)

type TConstants struct {
	Constantname  string         `db:"constantname"`
	Constantvalue sql.NullString `db:"constantvalue"`
}
