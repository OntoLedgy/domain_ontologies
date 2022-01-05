package dto

type EditEvent struct {
	Edit  int `db:"edit"`
	Event int `db:"event"`
}
