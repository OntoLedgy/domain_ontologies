package dto

type LabelUrl struct {
	ID      int    `db:"id"`
	LabelID int    `db:"label_id"`
	URL     string `db:"url"`
}
