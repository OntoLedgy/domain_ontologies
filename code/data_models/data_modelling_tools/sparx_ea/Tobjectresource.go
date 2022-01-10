package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TObjectresource struct {
	ObjectID        int             `db:"object_id"`
	Resource        string          `db:"resource"`
	Role            string          `db:"role"`
	Time            sql.NullFloat64 `db:"time"`
	Notes           sql.NullString  `db:"notes"`
	Percentcomplete sql.NullInt64   `db:"percentcomplete"`
	Datestart       pg.NullTime     `db:"datestart"`
	Dateend         pg.NullTime     `db:"dateend"`
	History         sql.NullString  `db:"history"`
	Expectedhours   sql.NullInt64   `db:"expectedhours"`
	Actualhours     sql.NullInt64   `db:"actualhours"`
}
