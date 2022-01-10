package sparx_ea

import (
	"database/sql"
)

type TTemplate struct {
	TemplateID   string         `db:"templateid"`
	Templatetype string         `db:"templatetype"`
	Templatename string         `db:"templatename"`
	Notes        sql.NullString `db:"notes"`
	Style        sql.NullString `db:"style"`
	Template     sql.NullString `db:"template"`
}
