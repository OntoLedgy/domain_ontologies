package dto

type SeriesTagRaw struct {
	Series   int  `db:"series"`
	Editor   int  `db:"editor"`
	Tag      int  `db:"tag"`
	IsUpvote bool `db:"is_upvote"`
}
