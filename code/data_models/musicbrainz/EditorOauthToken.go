package dto

import (
	"database/sql"
	"time"
)

type EditorOauthToken struct {
	ID                int            `db:"id"`
	Editor            int            `db:"editor"`
	Application       int            `db:"application"`
	AuthorizationCode sql.NullString `db:"authorization_code"`
	RefreshToken      sql.NullString `db:"refresh_token"`
	AccessToken       sql.NullString `db:"access_token"`
	ExpireTime        time.Time      `db:"expire_time"`
	Scope             int            `db:"scope"`
	Granted           time.Time      `db:"granted"`
}
