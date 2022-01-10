package sparx_ea

import (
	"database/sql"
)

type TConnectortypes struct {
	ConnectorType string         `db:"connector_type"`
	Description   sql.NullString `db:"description"`
}
