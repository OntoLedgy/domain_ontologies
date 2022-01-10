package sparx_ea

import (
	"database/sql"
)

type TSecuser struct {
	UserID     string         `db:"userid"`
	Userlogin  string         `db:"userlogin"`
	Firstname  string         `db:"firstname"`
	Surname    string         `db:"surname"`
	Department sql.NullString `db:"department"`
	Password   sql.NullString `db:"password"`
}
