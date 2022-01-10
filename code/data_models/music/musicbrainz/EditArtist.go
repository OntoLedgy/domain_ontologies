package dto

type EditArtist struct {
	Edit   int `db:"edit"`
	Artist int `db:"artist"`
	Status int `db:"status"`
}
