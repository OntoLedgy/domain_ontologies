package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TDiagram struct {
	DiagramID           int            `db:"diagram_id"`
	PackageID           sql.NullInt64  `db:"package_id"`
	ParentID            sql.NullInt64  `db:"parentid"`
	DiagramType         sql.NullString `db:"diagram_type"`
	Name                sql.NullString `db:"name"`
	Version             sql.NullString `db:"version"`
	Author              sql.NullString `db:"author"`
	Showdetails         sql.NullInt64  `db:"showdetails"`
	Notes               sql.NullString `db:"notes"`
	Stereotype          sql.NullString `db:"stereotype"`
	Attpub              int            `db:"attpub"`
	Attpri              int            `db:"attpri"`
	Attpro              int            `db:"attpro"`
	Orientation         sql.NullString `db:"orientation"`
	Cx                  sql.NullInt64  `db:"cx"`
	Cy                  sql.NullInt64  `db:"cy"`
	Scale               sql.NullInt64  `db:"scale"`
	Createddate         pg.NullTime    `db:"createddate"`
	Modifieddate        pg.NullTime    `db:"modifieddate"`
	Htmlpath            sql.NullString `db:"htmlpath"`
	Showforeign         int            `db:"showforeign"`
	Showborder          int            `db:"showborder"`
	Showpackagecontents int            `db:"showpackagecontents"`
	Pdata               sql.NullString `db:"pdata"`
	Locked              int            `db:"locked"`
	EaGuID              sql.NullString `db:"ea_guid"`
	Tpos                sql.NullInt64  `db:"tpos"`
	Swimlanes           sql.NullString `db:"swimlanes"`
	Styleex             sql.NullString `db:"styleex"`
}
