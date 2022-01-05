package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type ReplicationControl struct {
	ID                         int           `db:"id"`
	CurrentSchemaSequence      int           `db:"current_schema_sequence"`
	CurrentReplicationSequence sql.NullInt64 `db:"current_replication_sequence"`
	LastReplicationDate        pg.NullTime   `db:"last_replication_date"`
}
