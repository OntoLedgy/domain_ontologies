package dto

type LabelTagRaw struct {
	Label    int  `db:"label"`
	Editor   int  `db:"editor"`
	Tag      int  `db:"tag"`
	IsUpvote bool `db:"is_upvote"`
}
