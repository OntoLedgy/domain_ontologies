package dto

type InstrumentTagRaw struct {
	Instrument int  `db:"instrument"`
	Editor     int  `db:"editor"`
	Tag        int  `db:"tag"`
	IsUpvote   bool `db:"is_upvote"`
}
