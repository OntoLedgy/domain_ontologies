package dto

import (
	pg "github.com/lib/pq"
)

type EditNote struct {
	ID       int         `db:"id"`
	Editor   int         `db:"editor"`
	Edit     int         `db:"edit"`
	Text     string      `db:"text"`
	PostTime pg.NullTime `db:"post_time"`
}
