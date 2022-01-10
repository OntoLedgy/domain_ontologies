package sparx_ea

import (
	"database/sql"
)

type TClients struct {
	Name         string         `db:"name"`
	Organisation sql.NullString `db:"organisation"`
	Phone1       sql.NullString `db:"phone1"`
	Phone2       sql.NullString `db:"phone2"`
	Mobile       sql.NullString `db:"mobile"`
	Fax          sql.NullString `db:"fax"`
	Email        sql.NullString `db:"email"`
	Roles        sql.NullString `db:"roles"`
	Notes        sql.NullString `db:"notes"`
}
