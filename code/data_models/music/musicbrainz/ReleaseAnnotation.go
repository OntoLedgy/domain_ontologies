package dto

type ReleaseAnnotation struct {
	Release    int `db:"release"`
	Annotation int `db:"annotation"`
}
