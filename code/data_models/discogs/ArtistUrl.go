package dto

type ArtistUrl struct {
	ID       int    `db:"id"`
	ArtistID int    `db:"artist_id"`
	URL      string `db:"url"`
}
