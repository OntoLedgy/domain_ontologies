package sparx_ea

import (
	"database/sql"
)

type TObjectrisks struct {
	ObjectID int             `db:"object_id"`
	Risk     string          `db:"risk"`
	Risktype sql.NullString  `db:"risktype"`
	Evalue   sql.NullFloat64 `db:"evalue"`
	Notes    sql.NullString  `db:"notes"`
}
