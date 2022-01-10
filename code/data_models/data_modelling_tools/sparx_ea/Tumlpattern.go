package sparx_ea

import (
	"database/sql"
)

type TUmlpattern struct {
	PatternID       int            `db:"patternid"`
	Patterncategory sql.NullString `db:"patterncategory"`
	Patternname     sql.NullString `db:"patternname"`
	Style           sql.NullString `db:"style"`
	Notes           sql.NullString `db:"notes"`
	PatternXML      sql.NullString `db:"patternxml"`
	Version         sql.NullString `db:"version"`
}
