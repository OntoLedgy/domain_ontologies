package dto

type EditLabel struct {
	Edit   int `db:"edit"`
	Label  int `db:"label"`
	Status int `db:"status"`
}
