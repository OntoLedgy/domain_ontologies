package dto

type Tag struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	RefCount int    `db:"ref_count"`
}
