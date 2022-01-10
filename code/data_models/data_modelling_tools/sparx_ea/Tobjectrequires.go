package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TObjectrequires struct {
	ReqID       int            `db:"reqid"`
	ObjectID    sql.NullInt64  `db:"object_id"`
	Requirement sql.NullString `db:"requirement"`
	Reqtype     sql.NullString `db:"reqtype"`
	Status      sql.NullString `db:"status"`
	Notes       sql.NullString `db:"notes"`
	Stability   sql.NullString `db:"stability"`
	Difficulty  sql.NullString `db:"difficulty"`
	Priority    sql.NullString `db:"priority"`
	Lastupdate  pg.NullTime    `db:"lastupdate"`
}
