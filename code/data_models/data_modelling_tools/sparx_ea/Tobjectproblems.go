package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TObjectproblems struct {
	ObjectID      int            `db:"object_id"`
	Problem       string         `db:"problem"`
	Problemtype   string         `db:"problemtype"`
	Datereported  pg.NullTime    `db:"datereported"`
	Status        sql.NullString `db:"status"`
	Problemnotes  sql.NullString `db:"problemnotes"`
	Reportedby    sql.NullString `db:"reportedby"`
	Resolvedby    sql.NullString `db:"resolvedby"`
	Dateresolved  pg.NullTime    `db:"dateresolved"`
	Version       sql.NullString `db:"version"`
	Resolvernotes sql.NullString `db:"resolvernotes"`
	Priority      sql.NullString `db:"priority"`
	Severity      sql.NullString `db:"severity"`
}
