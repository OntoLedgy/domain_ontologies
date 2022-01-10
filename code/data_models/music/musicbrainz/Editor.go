package dto

import (
	"database/sql"

	pg "github.com/lib/pq"
)

type Editor struct {
	ID               int            `db:"id"`
	Name             string         `db:"name"`
	Privs            sql.NullInt64  `db:"privs"`
	Email            sql.NullString `db:"email"`
	Website          sql.NullString `db:"website"`
	Bio              sql.NullString `db:"bio"`
	MemberSince      pg.NullTime    `db:"member_since"`
	EmailConfirmDate pg.NullTime    `db:"email_confirm_date"`
	LastLoginDate    pg.NullTime    `db:"last_login_date"`
	LastUpdated      pg.NullTime    `db:"last_updated"`
	BirthDate        pg.NullTime    `db:"birth_date"`
	Gender           sql.NullInt64  `db:"gender"`
	Area             sql.NullInt64  `db:"area"`
	Password         string         `db:"password"`
	Ha1              string         `db:"ha1"`
	Deleted          bool           `db:"deleted"`
}
