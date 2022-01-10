package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TFiles struct {
	FileID    string         `db:"fileid"`
	Appliesto string         `db:"appliesto"`
	Category  string         `db:"category"`
	Name      string         `db:"name"`
	File      sql.NullString `db:"file"`
	Notes     sql.NullString `db:"notes"`
	Filedate  pg.NullTime    `db:"filedate"`
	Filesize  sql.NullInt64  `db:"filesize"`
}
