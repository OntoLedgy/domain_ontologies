package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TObject struct {
	ObjectID       int            `db:"object_id"`
	ObjectType     sql.NullString `db:"object_type"`
	DiagramID      sql.NullInt64  `db:"diagram_id"`
	Name           sql.NullString `db:"name"`
	Alias          sql.NullString `db:"alias"`
	Author         sql.NullString `db:"author"`
	Version        sql.NullString `db:"version"`
	Note           sql.NullString `db:"note"`
	PackageID      sql.NullInt64  `db:"package_id"`
	Stereotype     sql.NullString `db:"stereotype"`
	Ntype          sql.NullInt64  `db:"ntype"`
	Complexity     sql.NullString `db:"complexity"`
	Effort         sql.NullInt64  `db:"effort"`
	Style          sql.NullString `db:"style"`
	Backcolor      sql.NullInt64  `db:"backcolor"`
	Borderstyle    sql.NullInt64  `db:"borderstyle"`
	BorderwIDth    sql.NullInt64  `db:"borderwidth"`
	Fontcolor      sql.NullInt64  `db:"fontcolor"`
	Bordercolor    sql.NullInt64  `db:"bordercolor"`
	Createddate    pg.NullTime    `db:"createddate"`
	Modifieddate   pg.NullTime    `db:"modifieddate"`
	Status         sql.NullString `db:"status"`
	Abstract       sql.NullString `db:"abstract"`
	Tagged         sql.NullInt64  `db:"tagged"`
	Pdata1         sql.NullString `db:"pdata1"`
	Pdata2         sql.NullString `db:"pdata2"`
	Pdata3         sql.NullString `db:"pdata3"`
	Pdata4         sql.NullString `db:"pdata4"`
	Pdata5         sql.NullString `db:"pdata5"`
	Concurrency    sql.NullString `db:"concurrency"`
	Visibility     sql.NullString `db:"visibility"`
	Persistence    sql.NullString `db:"persistence"`
	Cardinality    sql.NullString `db:"cardinality"`
	Gentype        sql.NullString `db:"gentype"`
	Genfile        sql.NullString `db:"genfile"`
	Header1        sql.NullString `db:"header1"`
	Header2        sql.NullString `db:"header2"`
	Phase          sql.NullString `db:"phase"`
	Scope          sql.NullString `db:"scope"`
	Genoption      sql.NullString `db:"genoption"`
	Genlinks       sql.NullString `db:"genlinks"`
	Classifier     sql.NullInt64  `db:"classifier"`
	EaGuID         sql.NullString `db:"ea_guid"`
	ParentID       sql.NullInt64  `db:"parentid"`
	Runstate       sql.NullString `db:"runstate"`
	ClassifierGuID sql.NullString `db:"classifier_guid"`
	Tpos           sql.NullInt64  `db:"tpos"`
	Isroot         sql.NullInt64  `db:"isroot"`
	Isleaf         sql.NullInt64  `db:"isleaf"`
	Isspec         sql.NullInt64  `db:"isspec"`
	Isactive       sql.NullInt64  `db:"isactive"`
	Stateflags     sql.NullString `db:"stateflags"`
	Packageflags   sql.NullString `db:"packageflags"`
	Multiplicity   sql.NullString `db:"multiplicity"`
	Styleex        sql.NullString `db:"styleex"`
	Actionflags    sql.NullString `db:"actionflags"`
	Eventflags     sql.NullString `db:"eventflags"`
}
