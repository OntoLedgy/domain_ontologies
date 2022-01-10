package dto

type ReleaseTagRaw struct {
	Release  int  `db:"release"`
	Editor   int  `db:"editor"`
	Tag      int  `db:"tag"`
	IsUpvote bool `db:"is_upvote"`
}
