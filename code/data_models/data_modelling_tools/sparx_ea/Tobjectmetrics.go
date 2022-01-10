package sparx_ea

import (
	"database/sql"
)

type TObjectmetrics struct {
	ObjectID   int             `db:"object_id"`
	Metric     string          `db:"metric"`
	Metrictype sql.NullString  `db:"metrictype"`
	Evalue     sql.NullFloat64 `db:"evalue"`
	Notes      sql.NullString  `db:"notes"`
}
