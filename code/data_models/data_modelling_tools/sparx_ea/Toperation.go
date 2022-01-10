package sparx_ea

import (
	"database/sql"
)

type TOperation struct {
	OperationID  int            `db:"operationid"`
	ObjectID     sql.NullInt64  `db:"object_id"`
	Name         sql.NullString `db:"name"`
	Scope        sql.NullString `db:"scope"`
	Type         sql.NullString `db:"type"`
	Returnarray  sql.NullString `db:"returnarray"`
	Stereotype   sql.NullString `db:"stereotype"`
	Isstatic     sql.NullString `db:"isstatic"`
	Concurrency  sql.NullString `db:"concurrency"`
	Notes        sql.NullString `db:"notes"`
	Behaviour    sql.NullString `db:"behaviour"`
	Abstract     sql.NullString `db:"abstract"`
	Genoption    sql.NullString `db:"genoption"`
	Synchronized sql.NullString `db:"synchronized"`
	Pos          sql.NullInt64  `db:"pos"`
	Const        sql.NullInt64  `db:"const"`
	Style        sql.NullString `db:"style"`
	Pure         int            `db:"pure"`
	Throws       sql.NullString `db:"throws"`
	Classifier   sql.NullString `db:"classifier"`
	Code         sql.NullString `db:"code"`
	Isroot       sql.NullInt64  `db:"isroot"`
	Isleaf       sql.NullInt64  `db:"isleaf"`
	Isquery      sql.NullInt64  `db:"isquery"`
	Stateflags   sql.NullString `db:"stateflags"`
	EaGuID       sql.NullString `db:"ea_guid"`
	Styleex      sql.NullString `db:"styleex"`
}
