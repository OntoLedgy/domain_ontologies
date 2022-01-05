package dto

type WorkTagRaw struct {
	Work     int  `db:"work"`
	Editor   int  `db:"editor"`
	Tag      int  `db:"tag"`
	IsUpvote bool `db:"is_upvote"`
}
