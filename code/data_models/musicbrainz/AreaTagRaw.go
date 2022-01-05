package dto

type AreaTagRaw struct {
	Area     int  `db:"area"`
	Editor   int  `db:"editor"`
	Tag      int  `db:"tag"`
	IsUpvote bool `db:"is_upvote"`
}
