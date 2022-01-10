package dto

type ReleaseGroupAnnotation struct {
	ReleaseGroup int `db:"release_group"`
	Annotation   int `db:"annotation"`
}
