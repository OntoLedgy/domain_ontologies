package dto

type EventTagRaw struct {
	Event    int  `db:"event"`
	Editor   int  `db:"editor"`
	Tag      int  `db:"tag"`
	IsUpvote bool `db:"is_upvote"`
}
