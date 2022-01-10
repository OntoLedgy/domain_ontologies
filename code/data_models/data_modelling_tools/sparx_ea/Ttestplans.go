package sparx_ea

import (
	"database/sql"
)

type TTestplans struct {
	PlanID   string         `db:"planid"`
	Category sql.NullString `db:"category"`
	Name     string         `db:"name"`
	Author   sql.NullString `db:"author"`
	Notes    sql.NullString `db:"notes"`
	Testplan string         `db:"testplan"`
}
