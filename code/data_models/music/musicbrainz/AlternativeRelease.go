package dto

import (
	"database/sql"
)

type AlternativeRelease struct {
	ID           int            `db:"id"`
	GID          sql.NullString `db:"gid"`
	Release      int            `db:"release"`
	Name         sql.NullString `db:"name"`
	ArtistCredit sql.NullInt64  `db:"artist_credit"`
	Type         int            `db:"type"`
	Language     int            `db:"language"`
	Script       int            `db:"script"`
	Comment      string         `db:"comment"`
}
