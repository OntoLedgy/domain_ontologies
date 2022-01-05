package dto

type RecordingTagRaw struct {
	Recording int  `db:"recording"`
	Editor    int  `db:"editor"`
	Tag       int  `db:"tag"`
	IsUpvote  bool `db:"is_upvote"`
}
