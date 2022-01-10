package dto

type WorkRatingRaw struct {
	Work   int `db:"work"`
	Editor int `db:"editor"`
	Rating int `db:"rating"`
}
