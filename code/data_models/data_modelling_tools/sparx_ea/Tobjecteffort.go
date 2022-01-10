package sparx_ea

import (
	"database/sql"
)

type TObjecteffort struct {
	ObjectID   int             `db:"object_id"`
	Effort     string          `db:"effort"`
	Efforttype sql.NullString  `db:"efforttype"`
	Evalue     sql.NullFloat64 `db:"evalue"`
	Notes      sql.NullString  `db:"notes"`
}
