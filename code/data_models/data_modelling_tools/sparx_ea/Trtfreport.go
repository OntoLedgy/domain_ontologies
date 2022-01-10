package sparx_ea

import (
	"database/sql"
)

type TRtfreport struct {
	TemplateID      string         `db:"templateid"`
	Rootpackage     sql.NullInt64  `db:"rootpackage"`
	Filename        sql.NullString `db:"filename"`
	Details         sql.NullInt64  `db:"details"`
	Processchildren sql.NullInt64  `db:"processchildren"`
	Showdiagrams    sql.NullInt64  `db:"showdiagrams"`
	Heading         sql.NullString `db:"heading"`
	Requirements    sql.NullInt64  `db:"requirements"`
	Associations    sql.NullInt64  `db:"associations"`
	Scenarios       sql.NullInt64  `db:"scenarios"`
	Childdiagrams   sql.NullInt64  `db:"childdiagrams"`
	Attributes      sql.NullInt64  `db:"attributes"`
	Methods         sql.NullInt64  `db:"methods"`
	Imagetype       sql.NullInt64  `db:"imagetype"`
	Paging          sql.NullInt64  `db:"paging"`
	Intro           sql.NullString `db:"intro"`
	Resources       sql.NullInt64  `db:"resources"`
	Constraints     sql.NullInt64  `db:"constraints"`
	Tagged          sql.NullInt64  `db:"tagged"`
	Showtag         sql.NullInt64  `db:"showtag"`
	Showalias       sql.NullInt64  `db:"showalias"`
	Pdata1          sql.NullString `db:"pdata1"`
	Pdata2          sql.NullString `db:"pdata2"`
	Pdata3          sql.NullString `db:"pdata3"`
	Pdata4          sql.NullString `db:"pdata4"`
}
