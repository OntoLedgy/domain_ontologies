package dto

type PlaceTagRaw struct {
	Place    int  `db:"place"`
	Editor   int  `db:"editor"`
	Tag      int  `db:"tag"`
	IsUpvote bool `db:"is_upvote"`
}
