package sparx_ea

import (
	"database/sql"
)

type TTestclass struct {
	Testclass   string         `db:"testclass"`
	Description sql.NullString `db:"description"`
}
