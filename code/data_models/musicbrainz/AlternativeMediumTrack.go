package dto

type AlternativeMediumTrack struct {
	AlternativeMedium int `db:"alternative_medium"`
	Track             int `db:"track"`
	AlternativeTrack  int `db:"alternative_track"`
}
