package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TPhase struct {
	PhaseID      string         `db:"phaseid"`
	Phasename    string         `db:"phasename"`
	Phasenotes   sql.NullString `db:"phasenotes"`
	Phasestyle   sql.NullString `db:"phasestyle"`
	Startdate    pg.NullTime    `db:"startdate"`
	Enddate      pg.NullTime    `db:"enddate"`
	Phasecontent sql.NullString `db:"phasecontent"`
}
