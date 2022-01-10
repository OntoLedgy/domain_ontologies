package dto

import (
	"time"
)

type AutoeditorElectionVote struct {
	ID                 int       `db:"id"`
	AutoeditorElection int       `db:"autoeditor_election"`
	Voter              int       `db:"voter"`
	Vote               int       `db:"vote"`
	VoteTime           time.Time `db:"vote_time"`
}
