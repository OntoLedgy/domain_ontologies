package sparx_ea

import (
	"database/sql"
)

type TRtf struct {
	Type     sql.NullString `db:"type"`
	Template sql.NullString `db:"template"`
}
