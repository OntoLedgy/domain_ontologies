package dto

import (
	pg "github.com/lib/pq"
)

type Vote struct {
	ID         int         `db:"id"`
	Editor     int         `db:"editor"`
	Edit       int         `db:"edit"`
	Vote       int         `db:"vote"`
	VoteTime   pg.NullTime `db:"vote_time"`
	Superseded bool        `db:"superseded"`
}
