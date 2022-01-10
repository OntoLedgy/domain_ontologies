package dto

import (
	"database/sql"
	"time"
)

type EditorWatchPreferences struct {
	Editor                int            `db:"editor"`
	NotifyViaEmail        bool           `db:"notify_via_email"`
	NotificationTimeframe sql.NullString `db:"notification_timeframe"`
	LastChecked           time.Time      `db:"last_checked"`
}
