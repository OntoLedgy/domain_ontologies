package dto

type ArtistTagRaw struct {
	Artist   int  `db:"artist"`
	Editor   int  `db:"editor"`
	Tag      int  `db:"tag"`
	IsUpvote bool `db:"is_upvote"`
}
