package dto

type ReleaseGroupTagRaw struct {
	ReleaseGroup int  `db:"release_group"`
	Editor       int  `db:"editor"`
	Tag          int  `db:"tag"`
	IsUpvote     bool `db:"is_upvote"`
}
