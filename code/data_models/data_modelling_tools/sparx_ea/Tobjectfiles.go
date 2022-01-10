package sparx_ea

import (
	"database/sql"
)

type TObjectfiles struct {
	ObjectID int            `db:"object_id"`
	Filename string         `db:"filename"`
	Type     sql.NullString `db:"type"`
	Note     sql.NullString `db:"note"`
	Filesize sql.NullString `db:"filesize"`
	Filedate sql.NullString `db:"filedate"`
}
