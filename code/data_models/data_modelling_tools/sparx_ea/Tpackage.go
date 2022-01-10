package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TPackage struct {
	PackageID    int            `db:"package_id"`
	Name         sql.NullString `db:"name"`
	ParentID     sql.NullInt64  `db:"parent_id"`
	Createddate  pg.NullTime    `db:"createddate"`
	Modifieddate pg.NullTime    `db:"modifieddate"`
	Notes        sql.NullString `db:"notes"`
	EaGuID       sql.NullString `db:"ea_guid"`
	XMLpath      sql.NullString `db:"xmlpath"`
	Iscontrolled sql.NullInt64  `db:"iscontrolled"`
	Lastloaddate pg.NullTime    `db:"lastloaddate"`
	Lastsavedate pg.NullTime    `db:"lastsavedate"`
	Version      sql.NullString `db:"version"`
	Protected    sql.NullInt64  `db:"protected"`
	Pkgowner     sql.NullString `db:"pkgowner"`
	Umlversion   sql.NullString `db:"umlversion"`
	Usedtd       sql.NullInt64  `db:"usedtd"`
	LogXML       sql.NullInt64  `db:"logxml"`
	Codepath     sql.NullString `db:"codepath"`
	Namespace    sql.NullString `db:"namespace"`
	Tpos         sql.NullInt64  `db:"tpos"`
	Packageflags sql.NullString `db:"packageflags"`
	Batchsave    sql.NullInt64  `db:"batchsave"`
	Batchload    sql.NullInt64  `db:"batchload"`
}
