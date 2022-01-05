package dto

type LabelRatingRaw struct {
	Label  int `db:"label"`
	Editor int `db:"editor"`
	Rating int `db:"rating"`
}
