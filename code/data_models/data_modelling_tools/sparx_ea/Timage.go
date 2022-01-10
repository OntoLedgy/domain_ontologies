package sparx_ea

import (
	"database/sql"
)

type TImage struct {
	ImageID int            `db:"imageid"`
	Name    sql.NullString `db:"name"`
	Type    sql.NullString `db:"type"`
	Image   sql.NullString `db:"image"`
}
