package dto

type EventRatingRaw struct {
	Event  int `db:"event"`
	Editor int `db:"editor"`
	Rating int `db:"rating"`
}
