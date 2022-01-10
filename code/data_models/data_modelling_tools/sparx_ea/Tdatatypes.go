package sparx_ea

import (
	"database/sql"
)

type TDatatypes struct {
	Type         sql.NullString `db:"type"`
	Productname  sql.NullString `db:"productname"`
	Datatype     sql.NullString `db:"datatype"`
	Size         sql.NullInt64  `db:"size"`
	Maxlen       sql.NullInt64  `db:"maxlen"`
	Maxprec      sql.NullInt64  `db:"maxprec"`
	Maxscale     sql.NullInt64  `db:"maxscale"`
	Defaultlen   sql.NullInt64  `db:"defaultlen"`
	Defaultprec  sql.NullInt64  `db:"defaultprec"`
	Defaultscale sql.NullInt64  `db:"defaultscale"`
	User         sql.NullInt64  `db:"User"`
	Pdata1       sql.NullString `db:"pdata1"`
	Pdata2       sql.NullString `db:"pdata2"`
	Pdata3       sql.NullString `db:"pdata3"`
	Pdata4       sql.NullString `db:"pdata4"`
	Haslength    sql.NullString `db:"haslength"`
	Generictype  sql.NullString `db:"generictype"`
	DatatypeID   int            `db:"datatypeid"`
}
