package dto

type ArtistAnnotation struct {
	Artist     int `db:"artist"`
	Annotation int `db:"annotation"`
}
