package sparx_ea

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type TIssues struct {
	Issue        sql.NullString `db:"issue"`
	Issuedate    pg.NullTime    `db:"issuedate"`
	Owner        sql.NullString `db:"owner"`
	Status       sql.NullString `db:"status"`
	Notes        sql.NullString `db:"notes"`
	Resolver     sql.NullString `db:"resolver"`
	Dateresolved pg.NullTime    `db:"dateresolved"`
	Resolution   sql.NullString `db:"resolution"`
	IssueID      int            `db:"issueid"`
	Category     sql.NullString `db:"category"`
	Priority     sql.NullString `db:"priority"`
	Severity     sql.NullString `db:"severity"`
	Issuetype    sql.NullString `db:"issuetype"`
}
