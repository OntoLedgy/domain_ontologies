package dto

import (
	"time"
)

type ReleaseGroupSecondaryTypeJoin struct {
	ReleaseGroup  int       `db:"release_group"`
	SecondaryType int       `db:"secondary_type"`
	Created       time.Time `db:"created"`
}
