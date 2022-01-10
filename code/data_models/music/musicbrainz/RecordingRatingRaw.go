package dto

type RecordingRatingRaw struct {
	Recording int `db:"recording"`
	Editor    int `db:"editor"`
	Rating    int `db:"rating"`
}
