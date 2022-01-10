package dto

type ArtistNamevariation struct {
	ID       int    `db:"id"`
	ArtistID int    `db:"artist_id"`
	Name     string `db:"name"`
}
