package dto

import (
	"database/sql"
	"time"

	pg "github.com/lib/pq"
)

type AutoeditorElection struct {
	ID          int           `db:"id"`
	CandIDate   int           `db:"candidate"`
	Proposer    int           `db:"proposer"`
	Seconder1   sql.NullInt64 `db:"seconder_1"`
	Seconder2   sql.NullInt64 `db:"seconder_2"`
	Status      int           `db:"status"`
	YesVotes    int           `db:"yes_votes"`
	NoVotes     int           `db:"no_votes"`
	ProposeTime time.Time     `db:"propose_time"`
	OpenTime    pg.NullTime   `db:"open_time"`
	CloseTime   pg.NullTime   `db:"close_time"`
}
