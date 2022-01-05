package dto

type ArtistRatingRaw struct {
	Artist int `db:"artist"`
	Editor int `db:"editor"`
	Rating int `db:"rating"`
}
