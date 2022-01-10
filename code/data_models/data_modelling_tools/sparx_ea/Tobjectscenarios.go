package sparx_ea

import (
	"database/sql"
)

type TObjectscenarios struct {
	ObjectID     int             `db:"object_id"`
	Scenario     string          `db:"scenario"`
	Scenariotype sql.NullString  `db:"scenariotype"`
	Evalue       sql.NullFloat64 `db:"evalue"`
	Notes        sql.NullString  `db:"notes"`
	EaGuID       string          `db:"ea_guid"`
	XMLcontent   sql.NullString  `db:"xmlcontent"`
}
