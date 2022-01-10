package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TObjecttests struct {
	ObjectID           int            `db:"object_id"`
	Test               string         `db:"test"`
	Testclass          int            `db:"testclass"`
	Testtype           sql.NullString `db:"testtype"`
	Notes              sql.NullString `db:"notes"`
	Inputdata          sql.NullString `db:"inputdata"`
	Acceptancecriteria sql.NullString `db:"acceptancecriteria"`
	Status             sql.NullString `db:"status"`
	Daterun            pg.NullTime    `db:"daterun"`
	Results            sql.NullString `db:"results"`
	Runby              sql.NullString `db:"runby"`
	Checkby            sql.NullString `db:"checkby"`
}
