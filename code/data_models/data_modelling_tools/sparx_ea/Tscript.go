package sparx_ea

import (
	"database/sql"
)

type TScript struct {
	ScriptID       int            `db:"scriptid"`
	Scriptcategory sql.NullString `db:"scriptcategory"`
	Scriptname     sql.NullString `db:"scriptname"`
	Scriptauthor   sql.NullString `db:"scriptauthor"`
	Notes          sql.NullString `db:"notes"`
	Script         sql.NullString `db:"script"`
}
