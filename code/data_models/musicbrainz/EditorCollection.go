package dto

type EditorCollection struct {
	ID          int            `db:"id"`
	GID         sql.NullString `db:"gid"`
	Editor      int            `db:"editor"`
	Name        string         `db:"name"`
	Public      bool           `db:"public"`
	Description string         `db:"description"`
	Type        int            `db:"type"`
}
