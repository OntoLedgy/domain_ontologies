package dto

type RecordingAnnotation struct {
	Recording  int `db:"recording"`
	Annotation int `db:"annotation"`
}
