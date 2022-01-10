package dto

type ReleaseGroupRatingRaw struct {
	ReleaseGroup int `db:"release_group"`
	Editor       int `db:"editor"`
	Rating       int `db:"rating"`
}
