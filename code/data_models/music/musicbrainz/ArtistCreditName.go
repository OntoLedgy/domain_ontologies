package dto

type ArtistCreditName struct {
	ArtistCredit int    `db:"artist_credit"`
	Position     int    `db:"position"`
	Artist       int    `db:"artist"`
	Name         string `db:"name"`
	JoinPhrase   string `db:"join_phrase"`
}
