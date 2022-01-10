package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TTasks struct {
	TaskID     int            `db:"taskid"`
	Name       sql.NullString `db:"name"`
	Tasktype   sql.NullString `db:"tasktype"`
	Notes      sql.NullString `db:"notes"`
	Priority   sql.NullString `db:"priority"`
	Status     sql.NullString `db:"status"`
	Owner      sql.NullString `db:"owner"`
	Startdate  pg.NullTime    `db:"startdate"`
	Enddate    pg.NullTime    `db:"enddate"`
	Phase      sql.NullString `db:"phase"`
	History    sql.NullString `db:"history"`
	Percent    sql.NullInt64  `db:"percent"`
	Totaltime  sql.NullInt64  `db:"totaltime"`
	Actualtime sql.NullInt64  `db:"actualtime"`
	Assignedto sql.NullString `db:"assignedto"`
}
