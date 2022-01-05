package dto

type EditData struct {
	Edit int            `db:"edit"`
	Data sql.NullString `db:"data"`
}
