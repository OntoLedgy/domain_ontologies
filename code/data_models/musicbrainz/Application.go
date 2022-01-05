package dto

import (
	"database/sql"
)

type Application struct {
	ID               int            `db:"id"`
	Owner            int            `db:"owner"`
	Name             string         `db:"name"`
	OauthID          string         `db:"oauth_id"`
	OauthSecret      string         `db:"oauth_secret"`
	OauthRedirectUri sql.NullString `db:"oauth_redirect_uri"`
}
